<script lang="ts">
    import 'bootstrap/dist/css/bootstrap.min.css';
    import 'bootstrap/dist/js/bootstrap.min.js';
    import '@fortawesome/fontawesome-free/css/all.min.css'
    import './styles.scss';
    import {onMount} from "svelte";
    import {GetCwd} from "$lib/wailsjs/go/main/App";

    let cwd: string = "";

    onMount(() => {
        GetCwd().then((path) => {
            cwd = path
        })
    })
</script>

<div class="app">
    <main>
        <slot/>
    </main>

    <footer>
        <p class="text-center">
            App by <b>Pochie for Mochie</b>! @ All rights reserved <br>
            AppDir: {cwd}
        </p>
    </footer>
</div>

<style>
    .app {
        display: flex;
        flex-direction: column;
        min-height: 100vh;
    }

    main {
        flex: 1;
        display: flex;
        flex-direction: column;
        padding: 1rem;
        margin: 0 auto;
        box-sizing: border-box;
        width: 100%;
    }

    footer {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 12px;
        color: #999;
        font-style: italic;
        font-size: 12px;
    }

    @media (min-width: 480px) {
        footer {
            padding: 12px 0;
        }
    }
</style>
