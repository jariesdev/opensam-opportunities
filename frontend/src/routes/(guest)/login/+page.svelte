<script lang="ts">
    import logo from '../../../assets/images/logo.png'
    import LoginForm from "$lib/components/LoginForm.svelte";
    import { goto } from '$app/navigation'
    import type {User} from "$lib/types/user";

    let message: string = "Please enter your credentials 👇"
    let wasFailed: boolean = false

    function handleSuccess({detail: {isSuccess,username}}: CustomEvent): void {

        if (!isSuccess) {
            message = 'Invalid username or password!'
            wasFailed = true
            return
        }

        const user: User = {
            username: username
        }

        localStorage.setItem('auth_user', JSON.stringify(user))
        wasFailed = false
        goto('/opportunities')
    }
</script>

<svelte:head>
    <title>System Login</title>
</svelte:head>

<div class="container mt-5">
    <div class="login-wrap mx-auto">
        <img alt="App Logo" id="logo" class="img-fluid" src="{logo}">
        <div class="result" class:text-danger={wasFailed} id="result">
            {message}
        </div>
        <LoginForm on:success={handleSuccess}></LoginForm>
    </div>
</div>

<style lang="postcss">
    .login-wrap {
        text-align: center;
        max-width: 800px;
    }

    #logo {
        display: block;
        width: 50%;
        height: 50%;
        margin: auto;
        padding: 10% 0 0;
        background-position: center;
        background-repeat: no-repeat;
        background-size: 100% 100%;
        background-origin: content-box;
        max-width: 220px;
    }

    .result {
        height: 20px;
        line-height: 20px;
        margin: 1.5rem auto;
    }

</style>
