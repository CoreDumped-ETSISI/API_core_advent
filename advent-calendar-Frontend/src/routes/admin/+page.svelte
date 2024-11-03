<script>
    import { isAdmin } from '../../store'; // Import the isAdmin store
    import { onMount } from 'svelte'; // Optional for initializing states
    let password = '';
    let message = '';
    let loading = false; // State to manage loading

    const accessAdmin = async () => {
        loading = true; // Start loading state
        const response = await fetch('http://localhost:8000/admin/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ password }),
        });

        loading = false; // End loading state

        if (!response.ok) {
            const errorData = await response.json();
            message = `Error: ${errorData.error || 'Unknown error'}`;
            return;
        }

        // Successful access
        const data = await response.json();
        message = data.message || 'Access granted!';

        // Update the store to indicate the user is an admin
        isAdmin.set(true);
    };
</script>

<main>
    <h1>Admin Access</h1>
    <form on:submit|preventDefault={accessAdmin}>
        <input type="password" bind:value={password} placeholder="Enter admin password" required />
        <button type="submit" disabled={loading}>Access Admin</button>
    </form>
    {#if loading}
        <p>Loading...</p>
    {/if}
    <p>{message}</p>
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
    input {
        margin: 10px 0;
        padding: 10px;
        width: 100%;
        max-width: 300px; /* Limit width for better usability */
        border: 1px solid #ccc;
        border-radius: 4px;
    }
    button {
        padding: 10px 15px;
        background-color: #3498db;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
    }
    button:disabled {
        background-color: #ccc; /* Indicate loading state */
        cursor: not-allowed;
    }
    button:hover:not(:disabled) {
        background-color: #2980b9;
    }
</style>
