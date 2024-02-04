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
                <td>
                    {#if key === 'pointOfContact' && opportunity[key]}
                        {#each opportunity[key] as poc}
                            <ul class="list-unstyled">
                                <li><span class="fw-bold">Fax:</span> {poc.fax}</li>
                                <li><span class="fw-bold">Type:</span> {poc.type}</li>
                                <li><span class="fw-bold">Email:</span> {poc.email}</li>
                                <li><span class="fw-bold">Phone:</span> {poc.phone}</li>
                                <li><span class="fw-bold">Title:</span> {poc.title}</li>
                                <li><span class="fw-bold">FullName:</span> {poc.fullName}</li>
                            </ul>
                            {/each}
                    {:else if key === 'links' && opportunity[key]}
                        {#each opportunity[key] as link}
                            <ul class="list-unstyled">
                                <li>{link.href}</li>
                            </ul>
                        {/each}
                    {:else}
                        {opportunity[key]}
                    {/if}
                </td>
            </tr>
        {/each}
        </tbody>
    </table>
</div>