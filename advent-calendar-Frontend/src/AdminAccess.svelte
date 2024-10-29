<!-- src/AdminAccess.svelte -->
<script>
    import { isAdmin } from './store'; // Import the isAdmin store
    let password = '';
    let message = '';

    const accessAdmin = async () => {
        const response = await fetch('http://localhost:8000/admin/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ password }),
        });

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
    <input type="password" bind:value={password} placeholder="Enter admin password" />
    <button on:click={accessAdmin}>Access Admin</button>
    <p>{message}</p>
</main>
