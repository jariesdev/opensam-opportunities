<script lang="ts">
    import type {Header} from "$lib/components/dataTable/datatable";

    export let items: any[] = [];
    export let headers: Header[] = []
    export let loading: boolean = false

    function getColValue(header: Header, row: any): string {
        return row[header.field] || ''
    }
</script>

<div class="data-table">
    <div class="table-responsive">
        <table class="table">
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
                        <td scope={i === 0 ? 'row' : 'col'}>{getColValue(header, row)}</td>
                    {/each}
                </tr>
            {/each}
            </tbody>
        </table>
    </div>
    <div class="loading-indicator d-none" class:d-block={loading}>Loading...</div>
</div>