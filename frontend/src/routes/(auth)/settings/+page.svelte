<script lang="ts">
	import NavBar from "$lib/components/NavBar.svelte";
	import SettingsForm from "$lib/components/form/SettingsForm.svelte";
</script>

<svelte:head>
	<title>Settings</title>
</svelte:head>

<NavBar title="Settings"></NavBar>
<div class="text-column">
	<div class="container">
		<SettingsForm></SettingsForm>
	</div>
</div>
