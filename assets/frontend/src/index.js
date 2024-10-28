import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom/client";

const root = ReactDOM.createRoot(document.getElementById("root"));

const NavBar = () => {
  const [info, setInfo] = useState([]); 
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  useEffect(() => {
    const getData = async () => {
      try {
        let res = await fetch("http://localhost:8080/api");
        let data = await res.json();
        setInfo(data); 
      } catch (err) {
        console.error(err);
      }
    };

    getData(); 
  }, []);

  let items = info.map((el) => (
    <div key={el.id}>
      <span>{el.title}</span>
      <img src={el.image.split('/assets').join('')} width={150} alt={el.title} />
      <p>{el.content}</p>
    </div>
  ));

  return (
    <div>
      <nav>
        <label className="logo">KrakenGames</label>
        <ul>
          <li>
            <a href="#" className="active">Home</a>
          </li>
          <li
            className="dropdown"
            onMouseEnter={() => setIsDropdownOpen(true)}
            onMouseLeave={() => setIsDropdownOpen(false)}
          >
            <a href="#">Categories</a>
            {isDropdownOpen && (
              <ul className="dropdown-content" role="menu">
                <li role="menuitem"><a href="#" className="categorie">Action</a></li>
                <li role="menuitem"><a href="#" className="categorie">Adventure</a></li>
                <li role="menuitem"><a href="#" className="categorie">Puzzle</a></li>
                <li role="menuitem"><a href="#" className="categorie">Role-playing</a></li>
                <li role="menuitem"><a href="#" className="categorie">Simulation</a></li>
                <li role="menuitem"><a href="#" className="categorie">Strategy</a></li>
                <li role="menuitem"><a href="#" className="categorie">Sports</a></li>
              </ul>
            )}
          </li>
          <li><a href="#">Top Games</a></li>
          <li><a href="https://discord.gg/jtNZhDf9" target="_blank">Discord</a></li>
        </ul>
      </nav>
      <div>
        {items}
      </div>
    </div>
  );
};

root.render(<NavBar />);
