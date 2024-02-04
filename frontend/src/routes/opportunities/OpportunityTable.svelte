<script lang="ts">
    import { onMount } from 'svelte';
    import DataTable from "$lib/components/dataTable/DataTable.svelte";
    import { type Header} from "$lib/components/dataTable/datatable";
    import {PullLatest, SearchOpportunities} from "$lib/wailsjs/go/main/App";
    import Button from "$lib/components/button/Button.svelte";
    import Modal from "$lib/components/modal/Modal.svelte";
    import OpportunityDetails from "./OpportunityDetails.svelte";
    import {opportunity} from "$lib/wailsjs/go/models";
    import OpportunityFilter = opportunity.OpportunityFilter;
    import {debounce} from "lodash";

    let items: any[] = []
    let isLoading: boolean = false
    let isPulling: boolean = false
    let search: string = ""
    let showFilter: boolean = false

    let filters: OpportunityFilter = {
        fromDate: "",
        toDate: "",
        page: 1,
        perPage: 10
    }

    const headers: Header[] = [
        {
            label: 'Notice ID',
            field: 'noticeID'
        },
        {
            label: 'Title',
            field: 'title'
        },
        {
            label: 'Solicitation Number',
            field: 'solicitationNumber'
        },
        {
            label: 'Full Parent Path Name',
            field: 'fullParentPathName'
        },
        {
            label: 'Posted Date',
            field: 'postedDate'
        },
        {
            label: 'Type',
            field: 'type'
        },
        {
            label: 'UILink',
            field: 'uiLink'
        },
        {
            label: 'Actions',
            field: 'actions'
        }
    ]

    const loadData = debounce(function (): void {
        isLoading = true
        SearchOpportunities(search, filters)
            .then((opportunitiesData) => {
                items = opportunitiesData || []
            })
            .finally(() => {
                isLoading = false
            })
    }, 200, {maxWait: 1000})

    const pullLatest = debounce(function (): void {
        isPulling = true
        PullLatest()
            .then(() => {
                loadData()
            })
            .finally(() => {
                isPulling = false
            })
    }, 250)

    function toggleFilter() {
        showFilter = !showFilter
    }

    function downloadAsCsv() {
        // remove action field
        // headers.pop()
        // downloadCsvFromArray(items, headers, 'Export Report')
        SearchOpportunities(search, filters, true)
    }

    onMount(() => {
        loadData()
    })
</script>

<div class="opportunities-table">
    <div class="row mb-3">
        <div class="col-12 col-sm-4 col-md-3 mb-2 mb-sm-0">
            <div class="input-group">
                <input class="form-control" type="search" placeholder="Search" bind:value={search} on:input={loadData}>
                <span class="input-group-text">
                    <i class="fas fa-search"></i>
                </span>
            </div>
        </div>
        {#if showFilter}
            <div class="col-2">
                <input type="date" bind:value={filters.fromDate} class="form-control" placeholder="Posted From" on:input={loadData}>
            </div>
            <div class="col-2">
                <input type="date" bind:value={filters.toDate} class="form-control" placeholder="Posted To" on:input={loadData}>
            </div>
        {/if}
        <div class="col-12 col-sm-auto mx-auto">

        </div>
        <div class="col-12 col-sm-auto flex-grow-0 mb-2 mb-sm-0">
            <Button on:click={toggleFilter} class="text-nowrap w-100">
                {#if showFilter}
                    Hide Filters
                {:else}
                    Show Filters
                {/if}
            </Button>
        </div>
        <div class="col-12 col-sm-auto flex-grow-0 mb-2 mb-sm-0">
            <Button loading={isPulling} on:click={pullLatest} class="text-nowrap w-100">Pull Opportunities</Button>
        </div>
        <div class="col-12 col-sm-auto flex-grow-0 mb-2 mb-sm-0">
            <Button on:click={downloadAsCsv} class="text-nowrap w-100">Export</Button>
        </div>
    </div>
    <DataTable items={items} headers={headers} bind:page={filters.page} bind:perPage={filters.perPage} bind:loading={isLoading} on:filter={loadData}>
        <div slot="actions" let:row={row}>
            <Modal title="{row.title}">
                <div slot="activator" class="text-center text-nowrap" let:show={show}>
                    <a href="" on:click|preventDefault={show} >
                        <i class="fas fa-info-circle"></i>
                    </a>
                </div>
                <div slot="body">
                    <OpportunityDetails opportunity="{row}"></OpportunityDetails>
                </div>
            </Modal>
        </div>
    </DataTable>
</div>