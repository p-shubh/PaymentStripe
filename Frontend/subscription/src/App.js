// App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link, BrowserRouter } from 'react-router-dom';
import Home from './components/home';
import CardDetailsPage from './components/CardDetails';

const App = () => {
  return (
    <BrowserRouter>
    <Routes>
    <Route path="/" element={<Home/>} />
    <Route path="/card" element={<CardDetailsPage/>} />
    </Routes>
    </BrowserRouter>
  );
};

export default App;
