<script lang="ts">
	import OpportunityTable from "./OpportunityTable.svelte";
	import {goto} from "$app/navigation";

	function logout(): void {
		localStorage.setItem('auth_user', "")
		goto('/login')
	}
</script>

<svelte:head>
	<title>Opportunities</title>
</svelte:head>

<div class="text-column">
	<div class="container-fluid">
		<div class="d-flex justify-content-between">
			<h1 class="mb-4">Opportunities</h1>
			<button  type="button" class="btn btn-sm btn-primary-outline" on:click={logout}>logout</button>
		</div>
		<OpportunityTable></OpportunityTable>
	</div>
</div>
