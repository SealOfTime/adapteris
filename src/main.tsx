import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import { CssBaseline, StyledEngineProvider, ThemeProvider } from "@mui/material";

import theme from "themes";
import { BrowserRouter } from "react-router-dom";
import App from "./app";
import { QueryClientProvider } from "react-query";
import { queryClient } from "./react-query";

ReactDOM.render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <StyledEngineProvider injectFirst>
        <ThemeProvider theme={theme()}>
          <CssBaseline />
          <BrowserRouter>
            <App />
          </BrowserRouter>
        </ThemeProvider>
      </StyledEngineProvider>
    </QueryClientProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
