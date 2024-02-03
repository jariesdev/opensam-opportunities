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

    const headers: Header[] = [
        {
            label: 'Notice ID',
            field: 'noticeId'
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
        }
    ]

    function loadData(): void {
        isLoading = true
        SearchOpportunities(from, to)
            .then((result) => {
                const {opportunitiesData} = result
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

    onMount(() => {
        loadData()
    })
</script>

<div class="opportunities-table">
    <div class="d-flex justify-content-end mb-3">
            <Button bind:loading={isPulling} on:click={pullLatest}>Pull Opportunities</Button>
    </div>
    <DataTable items={items} headers={headers} bind:loading={isLoading}></DataTable>
</div>