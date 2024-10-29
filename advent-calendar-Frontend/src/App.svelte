<script>
    import { onMount } from 'svelte';
    import { Router, Route, Link } from 'svelte-routing';
    import AdminAccess from './AdminAccess.svelte';
    import Login from './Login.svelte';
    import Register from './Register.svelte';
    import CreateProblema from './CreateProblema.svelte';
    import Problema from './Problema.svelte';
    import Indice from './Indice.svelte';
    import NotFound from './Error.svelte';
    import { isAdmin } from './store';
    import {yearStore} from './store';

    let years = []; // Changed to a simple array
    let adminStatus = $isAdmin; // Get the current value of isAdmin store

    onMount(async () => {
        const response = await fetch('http://localhost:8000/'); // Update this URL as needed
        if (response.ok) {
            years = await response.json();
        } else {
            console.error("Failed to fetch years");
        }
    });
</script>

<Router>
    <nav>
        <Link to="/">Home</Link>
        <Link to="/admin">Admin</Link>
        <Link to="/login">Login</Link>
        <Link to="/register">Register</Link>
    </nav>

    <main>
        <h1>Welcome to My Svelte App</h1>
        
        <Route path="/" exact>
            <h2>This is the homepage.</h2>
            <h3>Select a Year to View Problems:</h3>
            <ul>
                {#each years.sort((a, b) => b - a) as year}
                    <li>
                        <Link to={`/problema/${year}`}>{year}</Link>
                    </li>
                {/each}
            </ul>
        </Route>

        <Route path="/admin" exact>
            {#if !adminStatus}
                <AdminAccess /> <!-- Show admin access form -->
            {:else}
                <div>
                    <h2>Admin Dashboard</h2>
                    <p>You have admin access!</p>
                    <Link to="/create-problema">Crear Problema</Link>
                    <button on:click={() => isAdmin.set(false)}>Log Out</button>
                </div>
            {/if}
        </Route>

        <Route 
            path="/problema/:year" 
            component={Indice} 
        />
        
        <Route path="/problema/:year/:day" component={Problema} exact />
        <Route path="/login" exact component={Login} />
        <Route path="/register" exact>
            <Register />
        </Route>
        <Route path="/create-problema" exact>
            {#if adminStatus}
                <CreateProblema />
            {:else}
                <h2>Debes ser administrador para acceder a esta página</h2>
            {/if}
        </Route>
        <Route path="*" exact>
            <NotFound />
        </Route>
    </main>
</Router>

<style>
    nav {
        display: flex;
        justify-content: space-around;
        margin: 20px;
    }
</style>
