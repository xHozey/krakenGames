import React, { useState, useEffect } from "react";
import "../index.css";
import GameCard from "../components/ui/GameCard";
import NavBar from "../components/ui/NavBar";

const HomePage = () => {
  const [info, setInfo] = useState([]);

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

  const items = info.map((el) => (
    <GameCard
      key={el.id}
      imageSrc={el.image.split("/assets").join("")}
      gameName={el.title}
    />
  ));

  return (
    <>
      <NavBar />
      <div className="blog-post">
        <div className="game-grid">{items}</div>
      </div>
    </>
  );
};

export default HomePage;
