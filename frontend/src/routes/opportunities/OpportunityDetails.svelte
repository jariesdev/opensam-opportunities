<script lang="ts">
    import {database} from "$lib/wailsjs/go/models";

    export let opportunity: database.Opportunity
</script>


<div>
    <table class="table table-sm table-hover">
        <thead>
            <tr>
                <th>Name</th>
                <th>Value</th>
            </tr>
        </thead>
        <tbody>
        {#each Object.keys(opportunity) as key}
            <tr>
                <td>{key}</td>
                <td>{opportunity[key]}</td>
            </tr>
        {/each}
        </tbody>
    </table>
</div>