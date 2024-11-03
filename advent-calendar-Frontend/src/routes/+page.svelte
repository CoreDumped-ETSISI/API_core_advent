<script>
    import { onMount } from 'svelte'; // Import the onMount lifecycle function

    // Declare a variable to store the fetched years
    let years = [];

    // Fetch years when the component is mounted
    onMount(async () => {
        try {
            const response = await fetch('http://localhost:8000/');
            if (response.ok) {
                years = await response.json();
            } else {
                console.error("Failed to fetch years");
            }
        } catch (error) {
            console.error("Error fetching years:", error); // Catch and log any errors
        }
    });
</script>

<nav>
    <a href="/">Home</a>
    <a href="/admin">Admin</a>
    <a href="/login">Login</a>
    <a href="/register">Register</a>
</nav>

<main>
    <h1>Welcome to My SvelteKit App</h1>
    
    <h2>This is the homepage.</h2>
    <h3>Select a Year to View Problems:</h3>
    <ul>
        {#if years.length > 0} <!-- Check if years has items -->
            {#each years.sort((a, b) => b - a) as year}
                <li>
                    <a href={`/problema/${year}`}>{year}</a>
                </li>
            {/each}
        {:else}
            <li>No years available.</li> <!-- Handle the case where there are no years -->
        {/if}
    </ul>
</main>

<style>
    nav {
        display: flex;
        justify-content: space-around;
        margin: 20px;
    }
</style>
