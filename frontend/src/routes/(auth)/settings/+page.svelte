<script lang="ts">
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
	<title>Settings</title>
</svelte:head>

<NavBar title="Settings"></NavBar>
<div class="text-column">
	<div class="container-fluid">

	</div>
</div>
