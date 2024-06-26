<script lang="ts">
    import { onMount } from 'svelte';
    import DataTable from "$lib/components/dataTable/DataTable.svelte";
    import { type Header} from "$lib/components/dataTable/datatable";
    import {
        GetOpportunityNaicsCodes,
        GetOpportunityTypes,
        PullLatest,
        SearchOpportunities
    } from "$lib/wailsjs/go/main/App";
    import Button from "$lib/components/button/Button.svelte";
    import Modal from "$lib/components/modal/Modal.svelte";
    import OpportunityDetails from "./OpportunityDetails.svelte";
    import {database, main, opportunity} from "$lib/wailsjs/go/models";
    import OpportunityFilter = opportunity.OpportunityFilter;
    import {debounce} from "lodash";
    import Multiselect from "$lib/components/input/Multiselect.svelte";
    import Opportunity = database.Opportunity;

    let items: Opportunity[] = []
    let total: number = 0
    let isLoading: boolean = false
    let isPulling: boolean = false
    let search: string = ""
    let showFilter: boolean = false
    let opportunityTypes: any[] = []
    let naicsCodes: any[] = []

    let filters: OpportunityFilter = {
        fromDate: "",
        toDate: "",
        type: [],
        naicsCode: [],
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
        SearchOpportunities(search, filters, false)
            .then((opportunitiesData) => {
                items = opportunitiesData.data || []
                total = opportunitiesData.total || 0
            })
            .finally(() => {
                isLoading = false
            })
    }, 200, {maxWait: 1000})

    const pullLatest = debounce(function (): void {
        isPulling = true
        PullLatest()
            .then((response: main.Response) => {
                if(!response.success) {
                    throw new Error(response.message)
                }
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

    function loadOptions() {
        GetOpportunityTypes().then((result) => {
            opportunityTypes = result || []
        })
        GetOpportunityNaicsCodes().then((result) => {
            naicsCodes = result || []
        })
    }

    function handleFilterChange() {
        filters.page = 1
        loadData()
    }

    function clearFilters(): void {
        filters.fromDate = ""
        filters.toDate = ""
        filters.type = []
        filters.naicsCode = []
        filters.page = 1
    }

    onMount(() => {
        loadData()
        loadOptions()
    })
</script>

<div class="opportunities-table">
    <div class="row mb-3">
        <div class="col-12 col-sm-4 col-md-3 mb-2 mb-sm-0">
            <div class="input-group">
                <input class="form-control" type="search" placeholder="Search" bind:value={search} on:input={handleFilterChange}>
                <span class="input-group-text">
                    {#if isLoading}
                        <i class="fas fa-spinner-third fa-spin"></i>
                    {:else}
                        <i class="fas fa-search"></i>
                    {/if}
                </span>
            </div>
        </div>
        {#if showFilter}
            <div class="col-12 col-sm-3 col-md-2 mb-2 mb-sm-0">
                <input type="date" bind:value={filters.fromDate} class="form-control" placeholder="Posted From" on:input={handleFilterChange}>
            </div>
            <div class="col-12 col-sm-3 col-md-2 mb-2 mb-sm-0">
                <input type="date" bind:value={filters.toDate} class="form-control" placeholder="Posted To" on:input={handleFilterChange}>
            </div>
            <div class="col-12 col-sm-3 col-md-2 mb-2 mb-sm-0">
                <Multiselect bind:value={filters.type} options="{opportunityTypes}" placeholder="Filter type" on:change={handleFilterChange} />
            </div>
            <div class="col-12 col-sm-3 col-md-2 mb-2 mb-sm-0">
                <Multiselect bind:value={filters.naicsCode} options="{naicsCodes}" placeholder="Filter code" on:change={handleFilterChange} />
            </div>
        {/if}
        <div class="col-12 {!showFilter ? 'col-sm-auto mx-auto' : ' mb-3'}">

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
            <Button loading={isPulling} on:click={pullLatest} class="text-nowrap w-100">Pull Latest</Button>
        </div>
        <div class="col-12 col-sm-auto flex-grow-0 mb-2 mb-sm-0">
            <Button on:click={downloadAsCsv} class="text-nowrap w-100">Export</Button>
        </div>
        {#if showFilter}
            <div class="col-12 col-sm-auto flex-grow-0 mb-2 mb-sm-0">
                <Button on:click={clearFilters} class="text-nowrap w-100">Clear Filter</Button>
            </div>
        {/if}
    </div>
    <DataTable items={items} headers={headers} bind:page={filters.page} bind:perPage={filters.perPage} total="{total}" bind:loading={isLoading} on:filter={loadData}>
        <div slot="actions" let:row={row}>
            <Modal title="{row.title}">
                <div slot="activator" class="text-center text-nowrap" let:show={show}>
                    <a href="#show-info" on:click|preventDefault={show} >
                        <i class="fas fa-info-circle"></i>
                    </a>
                </div>
                <div slot="body">
                    <OpportunityDetails opportunity="{row}" />
                </div>
            </Modal>
        </div>
    </DataTable>
</div>