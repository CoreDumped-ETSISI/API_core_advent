<script>
    import { getContext } from 'svelte';
    import { onMount } from 'svelte';
    import {yearStore} from './store';

    let year;
    let problemas = [];

    yearStore.subscribe(value => {
        year = value;
    });
    console.log(year);

    onMount(async () => {
        if (year) {
            try {
                const response = await fetch(`http://localhost:8000/${year}/`);
                if (response.ok) {
                    const data = await response.json();
                    problemas = data.sort((a, b) => a.dia - b.dia);
                } else {
                    console.error("Failed to fetch problems:", response.statusText);
                }
            } catch (error) {
                console.error("Error fetching problems:", error);
            }
        } else {
            console.error("Year parameter is missing in the URL.");

        }
    });

    export {problemas };


</script>

<main>

    <h1>Problemas de Advent of Code </h1>
    {#if problemas.length > 0}
        <ul>
            <p> A </p>
        </ul>
    {:else}
        <p>No hay problemas disponibles.</p>
    {/if}
</main>

<style>
    ul {
        list-style: none;
        padding: 0;
    }
    /* li {
        margin: 10px 0;
    }
    a {
        text-decoration: none;
        color: #007acc;
    }
    a:hover {
        text-decoration: underline;
    } */
</style>
