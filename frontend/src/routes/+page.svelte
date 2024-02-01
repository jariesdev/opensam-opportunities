<script lang="ts">
    import logo from '.././assets/images/logo.png'
    import LoginForm from "$lib/components/LoginForm.svelte";

    let message: string = "Please enter your credentials 👇"
    let wasFailed: boolean = false

    function handleSuccess({detail: {isSuccess}}: CustomEvent): void {

        if (!isSuccess) {
            message = 'Invalid username or password!'
            wasFailed = true
            return
        }

        wasFailed = false
        window.location.href = '/about'
    }
</script>

<svelte:head>
    <title>System Login</title>
</svelte:head>

<main>
    <div class="container">
        <div class="login-wrap">
            <img alt="App Logo" id="logo" class="img-fluid" src="{logo}">
            <div class="result" class:text-danger={wasFailed} id="result">
                {message}
            </div>
        </div>
    </div>
    <LoginForm on:success={handleSuccess}></LoginForm>
</main>

<style lang="postcss">
    .login-wrap {
        text-align: center;
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
