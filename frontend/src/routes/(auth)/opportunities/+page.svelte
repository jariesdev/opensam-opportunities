<script lang="ts">
	import OpportunityTable from "./OpportunityTable.svelte";
	import {goto} from "$app/navigation";
	import NavBar from "$lib/components/NavBar.svelte";

	function logout(): void {
		localStorage.setItem('auth_user', "")
		goto('/login')
	}
	function settings(): void {
		goto('/settings')
	}
</script>

<svelte:head>
	<title>Opportunities</title>
</svelte:head>

<NavBar title="Opportunities"></NavBar>
<div class="text-column">
	<div class="container-fluid">
		<OpportunityTable></OpportunityTable>
	</div>
</div>
