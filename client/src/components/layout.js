import React from "react";
// import "../scss/layout.scss";
import { Link } from "react-router-dom";
export function Layout() {
  return (
    <nav>        
        <ul>
          <Link to="/">
            <li>Home</li>
          </Link>
        </ul>     
    </nav>
  );
}