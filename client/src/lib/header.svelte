<nav class="header">
  <h2 class="header-logo">Reezan Visram</h2>

  <div class="header-links-container">
    <p class="header-link">Education</p>
    <p class="header-link">Experience</p>
    <p class="header-link">Projects</p>
    <button class="header-button">Resume</button>
  </div>
</nav>

<style>
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    margin-top: 1vw;
  }

  .header-logo {
    font-size: 32px;
    color: var(--text-primary-colour);
    margin-right: 25%;
  }

  .header-link {
    font-size: 24px;
    color: var(--text-primary-colour);
  }

  .header-links-container {
    display: flex;
    flex-grow: 1;
    justify-content: space-around;
    align-items: center;
  }

  .header-button {
    border-radius: 7.5px;
    outline: none;
    border: none;
    font-size: 24px;
    color: var(--button-text-colour);
    background-color: var(--text-primary-colour);
    width: 15%;
    padding-top: 1vh;
    padding-bottom: 1vh;
  }

  .header-button:hover {
    cursor: pointer;
  }
</style>