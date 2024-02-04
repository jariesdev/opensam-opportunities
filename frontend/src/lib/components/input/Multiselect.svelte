<script lang="ts">
    export let value: string[] = []
    export let options: any[] = []

    function getValue(option: any) {
        return (typeof option === 'string' ? option :  option['value']) || ""
    }

    function getText(option: any) {
        return (typeof option === 'string' ? option :  option['text']) || ""
    }
</script>


<div class="multi-select">
    <select class="form-control" bind:value={value} multiple on:change on:input>
        {#each options as option}
            <option value="{getValue(option)}">{getText(option)}</option>
        {/each}
    </select>
</div>