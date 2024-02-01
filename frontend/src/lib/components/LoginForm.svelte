<script lang="ts">
    import {Login} from '../wailsjs/go/main/App.js'
    import {createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher();

    let username: string
    let password: string
    let isProcessing: boolean = false
    let isSuccess: boolean = false

    let resultText: string

    function login(): void {
        isProcessing = true

        Login(username, password)
            .then(({Message, Result}) => {
                isSuccess = Result
                resultText = Message
                console.log(Message, Result)
                dispatch('success', {isSuccess,resultText})
            })
            .finally(() => {
                isProcessing = false
            })
    }
</script>

<section>
    <form on:submit|preventDefault={login}>
        <div class="input-box" id="input">
            <input autocomplete="off"
                   bind:value={username}
                   class="input"
                   id="name"
                   type="text"
                   placeholder="Username"/>
            <input autocomplete="off"
                   bind:value={password}
                   class="input"
                   id="password"
                   type="password"
                   placeholder="Password"/>
            <button type="submit"
                    class="btn"
                    disabled="{isProcessing || !username || !password}">
                Login
            </button>
        </div>
    </form>
</section>

<style>

    .input-box .btn {
        display: block;
        height: 30px;
        line-height: 30px;
        border-radius: 3px;
        border: none;
        margin: 0 auto 20px;
        padding: 0 30px;
        cursor: pointer;
    }

    .input-box .btn:hover {
        background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
        color: #333333;
    }

    .input-box .input {
        border: 1px solid rgba(240, 240, 240, 1);
        border-radius: 3px;
        outline: none;
        height: 30px;
        line-height: 30px;
        padding: 0 10px;
        background-color: rgba(240, 240, 240, 1);
        -webkit-font-smoothing: antialiased;
        width: 100%;
        max-width: 250px;
        display: block;
        transition: all 250ms ease-in;
        margin: 0 auto 0.75rem;
    }

    .input-box .input:focus {
        border-color: rgba(240, 240, 240, 1);
        background-color: rgba(255, 255, 255, 1);
    }

</style>
