<script lang="ts">
    import {goto} from "$app/navigation";

    export let title:string = ""

    function logout(): void {
        localStorage.setItem('auth_user', "")
        goto('/login')
    }
</script>

<div class="container-fluid">
    <div class="d-flex justify-content-between mb-5">
        <h1 class="mb-4">{title}</h1>
        <div class="d-flex justify-content-end">
            <button type="button" class="btn btn-sm btn-primary-outline" on:click={()=>goto('opportunities')}>
                <i class="fas fa-briefcase"></i>
                Opportunities
            </button>
            <button type="button" class="btn btn-sm btn-primary-outline" on:click={()=>goto('settings')}>
                <i class="fas fa-gear"></i>
                Settings
            </button>
            <button type="button" class="btn btn-sm btn-primary-outline" on:click={logout}>
                <i class="fas fa-sign-out"></i>
                Logout
            </button>
        </div>
    </div>
</div>