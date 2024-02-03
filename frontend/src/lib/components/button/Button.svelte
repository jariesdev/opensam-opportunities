<button class="btn btn-primary px-3 {$$restProps.class || ''} {`btn-${size}`}"
        type="button"
        disabled="{disabled || loading}"
        on:click>
    {#if loading}
        <span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
    {/if}
    <slot>
        Button Text
    </slot>
</button>


<script lang="ts">
    export let disabled: boolean = false
    export let loading: boolean = false
    export let size: string = false
</script>