import { Route, Routes } from 'solid-app-router';
import { Component, createEffect, createSignal, lazy } from 'solid-js';
import Nav from './components/nav';

const Home = lazy(() => import('./pages/home'));
const Saved = lazy(() => import('./pages/saved'));

const [username, setUserName] = createSignal('m3rashid');
const [repos, setRepos] = createSignal([]);

createEffect(async () => {
  const res = await fetch(`https://api.github.com/users/${username()}/repos`, {
    headers: {
      Authorization: 'token <TOKEN>',
      Accept: 'application/vnd.github+json',
    },
  });
  console.log(res);
  setRepos(await res.json());
});

const App: Component = () => {
  return (
    <div class='container'>
      <Nav />
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/saved' element={<Saved />} />
      </Routes>
    </div>
  );
};

export { username, setUserName, repos };
export default App;
