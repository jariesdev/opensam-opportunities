<script lang="ts">
    import filter from "lodash/filter";
    import {createEventDispatcher} from 'svelte'
    import {clickOutside} from "$lib/utils/clickOutside";

    export let value: string[] = []
    export let options: any[] = []
    export let placeholder: string = "select type"
    let showDropdown: boolean = false
    let searchInput: string = ""
    const dispatcher = createEventDispatcher()

    function getValue(option: any) {
        return (typeof option === 'string' ? option :  option['value']) || ""
    }

    function getText(option: any) {
        return (typeof option === 'string' ? option :  option['text']) || ""
    }

    function removeSelection(option:any) {
        value = filter(value, v=> v !== getValue(option))
    }

    function toggleDropdown() {
        showDropdown = !showDropdown
    }

    function handleClickOutside() {
        showDropdown = false
    }

    $: visibleOptions = () => {
        if(searchInput) {
            return filter(options, o => {
                let str = getText(o).toLowerCase()
                return str.includes(searchInput.toLowerCase())
            })
        }

        return options
    }

    $: selectionOptions = () => {
        return filter(options, o => value.includes(getValue(o)))
    }

    $: if (value) {
        dispatcher('input', value)
        dispatcher('change', value)
    }

</script>


<div use:clickOutside class="multi-select" on:click_outside={handleClickOutside}>
    <div class="m-dropdown position-relative">
        <div class="selected-options form-control grid" role="button" tabindex="0" on:click={toggleDropdown}>
            {#if selectionOptions().length === 0}
                <div class="btn border-0 btn-sm disabled py-0">
                    {placeholder}
                </div>
            {/if}
            {#each selectionOptions() as option}
                <button type="button" class="btn btn-primary btn-sm py-0 me-1" on:click|stopPropagation={removeSelection(option)}>
                    {#if getText(option)}
                        {getText(option)}
                    {:else}
                        <span>&nbsp;</span>
                    {/if}
                </button>
            {/each}
        </div>
        <div class="m-dropdown-menu z-10 position-absolute bg-white w-100 rounded-2 shadow rounded-bottom overflow-hidden pb-1 fade {showDropdown ? 'show d-block' : 'd-none'}">
            <div class="filter-input">
                <input bind:value="{searchInput}" type="search" class="form-control-plaintext px-2" placeholder="search">
            </div>
            <div class="option-wrap">
                <ul class="list-group">
                    {#each visibleOptions() as option}
                        <li class="list-group-item rounded-0">
                            <label class="form-check-label d-block">
                                <input bind:group={value} class="form-check-input me-1" type="checkbox" value="{getValue(option)}" aria-label="{getText(option)}">
                                {getText(option)}
                            </label>
                        </li>
                    {/each}
                </ul>
            </div>
        </div>
    </div>

<!--    <select class="form-control" bind:value={value} multiple on:change on:input>-->
<!--        {#each options as option}-->
<!--            <option value="{getValue(option)}">{getText(option)}</option>-->
<!--        {/each}-->
<!--    </select>-->
</div>

<style lang="scss">
    .option-wrap {
        max-height: 400px;
        overflow: auto;
    }
    .selected-options {
        .btn {

        }
    }
</style>