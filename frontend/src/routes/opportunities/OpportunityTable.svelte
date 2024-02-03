<script lang="ts">
    import { onMount } from 'svelte';
    import DataTable from "$lib/components/dataTable/DataTable.svelte";
    import { type Header} from "$lib/components/dataTable/datatable";
    import {PullLatest, SearchOpportunities} from "$lib/wailsjs/go/main/App";
    import Button from "$lib/components/button/Button.svelte";

    let items: any[] = []
    const from: string = '12/01/2023'
    const to: string = '12/07/2023'
    let isLoading: boolean = false
    let isPulling: boolean = false
    let search: string = ""
    let showFilter: boolean = false

    interface OpportunityFilter {
        fromDate: string
        toDate: string
    }

    let filters: OpportunityFilter = {
        fromDate: "",
        toDate: ""
    }

    const headers: Header[] = [
        {
            label: 'Notice ID',
            field: 'NoticeID'
        },
        {
            label: 'Title',
            field: 'Title'
        },
        {
            label: 'Solicitation Number',
            field: 'SolicitationNumber'
        },
        {
            label: 'Full Parent Path Name',
            field: 'FullParentPathName'
        },
        {
            label: 'Posted Date',
            field: 'PostedDate'
        },
        {
            label: 'Type',
            field: 'Type'
        }
    ]

    function loadData(): void {
        isLoading = true
        SearchOpportunities(search, filters)
            .then((opportunitiesData) => {
                console.log(opportunitiesData)
                items = opportunitiesData || []
            })
            .finally(() => {
                isLoading = false
            })
    }

    function pullLatest() {
        isPulling = true
        PullLatest().finally(() => {
            isPulling = false
        })
    }

    function toggleFilter() {
        showFilter = !showFilter
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
    </div>
    <DataTable items={items} headers={headers} bind:loading={isLoading}></DataTable>
</div>