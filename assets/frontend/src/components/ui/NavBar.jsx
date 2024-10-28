import React, { useState } from 'react';

function Navbar() {
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  return (
    <nav className="navbar">
      <label className="logo">KrakenGames</label>
      <ul className="nav-links">
        <li><a href="#" className="active">Home</a></li>
        <li 
          className="dropdown"
          onMouseEnter={() => setIsDropdownOpen(true)}
          onMouseLeave={() => setIsDropdownOpen(false)}
        >
          <a href="#">Categories</a>
          {isDropdownOpen && (
            <ul className="dropdown-content">
              <li><a href="#" className="categorie">Action</a></li>
              <li><a href="#" className="categorie">Adventure</a></li>
              <li><a href="#" className="categorie">Puzzle</a></li>
              <li><a href="#" className="categorie">Role-playing</a></li>
              <li><a href="#" className="categorie">Simulation</a></li>
              <li><a href="#" className="categorie">Strategy</a></li>
              <li><a href="#" className="categorie">Sports</a></li>
            </ul>
          )}
        </li>
        <li><a href="#">Top Games</a></li>
        <li><a href="https://discord.gg/jtNZhDf9" target="_blank" rel="noopener noreferrer">Discord</a></li>
      </ul>
    </nav>
  );
}

export default Navbar;