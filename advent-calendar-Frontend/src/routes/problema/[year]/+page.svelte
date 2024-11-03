<!-- src/routes/problems/[year]/+page.svelte -->
<script lang="ts">
    import type { PageData } from "../$types";
    export let data: PageData
    const { year, problemas, error } = data;

    if (error) {
        console.error(error);
    }
</script>

<main>
    <h1>Problemas de Advent of Code - Año {year}</h1>

    {#if problemas.length > 0}
        <ul>
            {#each problemas as problema}
                <li>
                    <a href={`/problema/${year}/${problema.dia}`}>Día {problema.dia}</a>
                </li>
            {/each}
        </ul>
    {:else}
        <p>Inicie sesión para ver los problemas.</p>
    {/if}
</main>

<style>
    main {
        text-align: center;
        padding: 20px;
    }
    h1 {
        font-size: 2rem;
        color: #3498db; /* Customize your color */
    }
    ul {
        list-style: none;
        padding: 0;
    }
    li {
        margin: 10px 0;
    }
    a {
        color: #2980b9;
        text-decoration: none;
    }
    a:hover {
        text-decoration: underline;
    }
</style>
