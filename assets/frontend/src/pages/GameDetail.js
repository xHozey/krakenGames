import React from 'react';

const GameDetails = ({ game }) => {
  return (
    <div className="game-details-container">
      <div className="game-header">
        <img src={game.image} alt={game.title} className="game-image" />
        <h1 className="game-title">{game.title}</h1>
      </div>
      <div className="game-content">
        <p className="game-description">{game.content}</p>
      </div>
      <div className="game-screenshots">
        <h2 className="section-title">Screenshots</h2>
        <div className="screenshot-grid">
          {game.screens.map((screen, index) => (
            <img key={index} src={screen} alt={`Screenshot ${index + 1}`} className="screenshot" />
          ))}
        </div>
      </div>
      <div className="game-info-section">
        <h2 className="section-title">Game Information</h2>
        <div className="info-grid">
          <div className="info-item">
            <span className="info-label">Genre:</span>
            <span className="info-value">{game.info.genre.join(', ')}</span>
          </div>
          <div className="info-item">
            <span className="info-label">Developer:</span>
            <span className="info-value">{game.info.developer}</span>
          </div>
          <div className="info-item">
            <span className="info-label">Platform:</span>
            <span className="info-value">{game.info.platform}</span>
          </div>
          <div className="info-item">
            <span className="info-label">Game Size:</span>
            <span className="info-value">{game.info.game_size}</span>
          </div>
          <div className="info-item">
            <span className="info-label">Release Date:</span>
            <span className="info-value">{game.info.released_date}</span>
          </div>
          <div className="info-item">
            <span className="info-label">Version:</span>
            <span className="info-value">{game.info.version}</span>
          </div>
        </div>
      </div>
      <div className="game-info-section">
        <h2 className="section-title">System Requirements</h2>
        <div className="info-grid">
          <div className="info-item">
            <span className="info-label">OS:</span>
            <span className="info-value">{game.system.os}</span>
          </div>
          <div className="info-item">
            <span className="info-label">Processor:</span>
            <span className="info-value">{game.system.processor}</span>
          </div>
          <div className="info-item">
            <span className="info-label">Memory:</span>
            <span className="info-value">{game.system.memory}</span>
          </div>
          <div className="info-item">
            <span className="info-label">Graphics:</span>
            <span className="info-value">{game.system.graphics}</span>
          </div>
          <div className="info-item">
            <span className="info-label">DirectX:</span>
            <span className="info-value">{game.system.direct_x}</span>
          </div>
          <div className="info-item">
            <span className="info-label">Storage:</span>
            <span className="info-value">{game.system.storage}</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default GameDetails;