<script>
    import { goto } from '$app/navigation'; // Use $app/navigation for navigation in SvelteKit
</script>

<main>
    <h1>404 - Page Not Found</h1>
    <p>Sorry, the page you are looking for does not exist.</p>
    <button on:click={() => goto('/')}>Go to Home</button>
</main>

<style>
    main {
        text-align: center;
        padding: 50px;
    }
    h1 {
        font-size: 2rem;
        color: #e74c3c; /* You can choose a color of your choice */
    }
    button {
        padding: 10px 15px;
        background-color: #3498db;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
    }
    button:hover {
        background-color: #2980b9;
    }
</style>
