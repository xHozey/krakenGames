import React, { useState } from 'react';
import PropTypes from 'prop-types';

function GameCard({ imageSrc, gameName }) {
  const [isHovered, setIsHovered] = useState(false);
  return (
    <div 
      className="game-card"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      <img src={imageSrc} alt={gameName} className="game-image" />
      {isHovered && (
        <div className="game-overlay">
          <h3 className="game-name">{gameName}</h3>
        </div>
      )}

    </div>
  );
}

GameCard.propTypes = {
  imageSrc: PropTypes.string.isRequired,
  gameName: PropTypes.string.isRequired
};

export default GameCard;