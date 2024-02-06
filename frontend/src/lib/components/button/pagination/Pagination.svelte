<script lang="ts">
    import {createEventDispatcher} from "svelte";

    export let value: number = 1
    export let total: number = 0
    export let size: number = 15
    export let maxPager: number = 5

    const dispatcher = createEventDispatcher()

    function setCurrentPage(page: number): void {
        value = page
        dispatcher('change', value)
    }

    $: numberOfPages = ():number => total > 0 ? Math.ceil(total / size) : 1
    $: pagers = () => {
        const p = []
        const median = Math.floor(maxPager / 2)
        let start: number;
        let end: number;
        if (value <= median) {
            start = 1;
            end = maxPager
        } else if(value >= numberOfPages() - median ) {
            start = numberOfPages() - (maxPager - 1)
            end = numberOfPages()
        }else {
            start = value - median;
            end = value + median
        }

        while (start <= end) {
            p.push(start++)
        }
        return p
    }


</script>

<nav aria-label="Pagination">
    <ul class="pagination mb-0">
        <li class="page-item">
            <a class="page-link" class:disabled={value === 1} href="#first" on:click="{() => setCurrentPage(1)}">
                <i class="fas fa-angle-double-left"></i>
            </a>
        </li>
        <li class="page-item">
            <a class="page-link" class:disabled={value === 1} href="#previous" on:click="{() => setCurrentPage(value - (value > 1 ? 1 : 0))}">
                <i class="fas fa-angle-left"></i>
            </a>
        </li>
        {#each pagers() as page}
            <li class="page-item">
                <a class="page-link" class:active={value===page} href="#goto-page-{page}" on:click="{() => setCurrentPage(page)}">
                    {page}
                </a>
            </li>
        {/each}
        <li class="page-item">
            <a class="page-link" class:disabled={value === numberOfPages()} href="#next" on:click="{() => setCurrentPage(value + (value < numberOfPages() ? 1 : 0))}">
                <i class="fas fa-angle-right"></i>
            </a>
        </li>
        <li class="page-item">
            <a class="page-link" class:disabled={value === numberOfPages()} href="#last" on:click="{() => setCurrentPage(numberOfPages())}">
                <i class="fas fa-angle-double-right"></i>
            </a>
        </li>
    </ul>
</nav>