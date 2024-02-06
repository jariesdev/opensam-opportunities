<script lang="ts">
    import loadingImg from '../../../../../frontend/src/assets/images/pikachu.gif'
    import type {Header} from "$lib/components/dataTable/datatable";
    import { createEventDispatcher } from 'svelte'
    import Pagination from "$lib/components/button/pagination/Pagination.svelte";

    export let items: any[] = [];
    export let headers: Header[] = []
    export let loading: boolean = false
    export let total: number = 0
    export let page: number = 1
    export let perPage: number = 20
    export let wordWrap: boolean = false
    let perPageOptions: number[] = [10,20,50,100,1000]

    const dispatch = createEventDispatcher()

    function getColValue(header: Header, row: any): string {
        return row[header.field] || ''
    }

    function handleFilterChange() {
        dispatch('filter')
    }
</script>

<div class="data-table">
    <div class="table-responsive">
        <table class="table {wordWrap ? 'table-sm single-line-content' : ''}">
            <thead>
            <tr>
                {#each headers as header}
                    <th scope="col">{header.label}</th>
                {/each}
            </tr>
            </thead>
            <tbody>
            {#if items.length === 0}
                <tr>
                    <td colspan="{headers.length}">
                        <p class="text-muted small text-center fst-italic">
                            No records yet.
                        </p>
                    </td>
                </tr>
                {/if}
            {#each items as row}
                <tr>
                    {#each headers as header, i}
                        <td>
                            <slot name="column" {row} {header} {getColValue}>
                                {getColValue(header, row)}
                            </slot>

                            {#if header.field === 'actions'}
                                <slot name="actions" {row} {getColValue}></slot>
                            {/if}
                        </td>
                    {/each}
                </tr>
            {/each}
            </tbody>
        </table>
    </div>
    <div class="row mt-3 ">
        <div class="col-2 align-self-center">
            <span class="fst-italic">{total} total records</span>
        </div>
        <div class="col mx-auto"></div>
        <div class="col flex-grow-0 text-nowrap align-self-center">
            <label class="form-check">
                <input type="checkbox" bind:checked="{wordWrap}" class="form-check-input">
                Compact View
            </label>
        </div>
        <div class="col-2">
            <select class="form-control" bind:value={perPage} on:change={handleFilterChange}>
                {#each perPageOptions as option}
                    <option value="{option}">{option} per page</option>
                {/each}
            </select>
        </div>
        <div class="col flex-grow-0">
            <Pagination bind:value={page} total="{total}" size="{perPage}" on:change={handleFilterChange} />
        </div>
    </div>
    <div class="loading-indicator text-nowrap me-3 mb-3 {!loading ? 'd-none' : ''}" >
        <img src="{loadingImg}" alt="" class="loading-pikachu">
        <span class="sr-only">Loading</span>
    </div>
</div>

<style lang="scss">
    .loading-indicator {
        pointer-events: none;
        position: fixed;
        bottom: 0;
        right: 0;
    }
    .loading-pikachu {
        width: 50px;
    }

    table {
        &.single-line-content {
            td {
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
                max-width: 300px;
            }
        }
    }
</style>