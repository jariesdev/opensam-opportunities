<script lang="ts">
    import {Login} from "$lib/wailsjs/go/main/App";
    import {createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher();

    let username: string = 'admin'
    let password: string = 'password'
    let isProcessing: boolean = false
    let isSuccess: boolean = false

    let resultText: string

    function login(): void {
        isProcessing = true

        Login(username, password)
            .then(({message, result}) => {
                isSuccess = result
                resultText = message

                dispatch('success', {isSuccess,resultText})
            })
            .finally(() => {
                isProcessing = false
            })
    }
</script>

<section>
    <div class="row">
        <div class="col-sm-5 mx-auto">
            <form on:submit|preventDefault={login}>
                <div class="input-box" id="input">
                    <input autocomplete="off"
                           bind:value={username}
                           class="form-control mb-3"
                           id="name"
                           type="text"
                           placeholder="Username"/>
                    <input autocomplete="off"
                           bind:value={password}
                           class="form-control mb-3"
                           id="password"
                           type="password"
                           placeholder="Password"/>
                    <button type="submit"
                            class="btn btn-primary btn-block d-block px-5 mx-auto"
                            disabled="{isProcessing || !username || !password}">
                        Login
                    </button>
                </div>
            </form>
        </div>
    </div>
</section>

