// App.jsx
import React, { useEffect, useState } from 'react';
import { fetchCats } from './api';

function App() {
  const [cats, setCats] = useState([]);

  useEffect(() => {
    const getCats = async () => {
      const catsData = await fetchCats();
      setCats(catsData);
    };
    getCats();
  }, []);

  return (
    <div>
      <h1>Cat Farm</h1>
      <ul>
        {cats.map(cat => (
          <li key={cat.id}>{cat.name} ({cat.breed})</li>
        ))}
      </ul>
    </div>
  );
}

export default App;