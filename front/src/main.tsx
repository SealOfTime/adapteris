import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import { CssBaseline, StyledEngineProvider, ThemeProvider } from "@mui/material";

import theme from "themes";
import { BrowserRouter } from "react-router-dom";
import App from "./app";
import { QueryClientProvider } from "react-query";
import { queryClient } from "./react-query";
import {LocalizationProvider} from "@mui/lab";
import {formatDate} from "hooks/useFormatDate";
import AdapterDateFns from '@mui/lab/AdapterDateFns';

ReactDOM.render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <StyledEngineProvider injectFirst>
        <ThemeProvider theme={theme()}>
          <LocalizationProvider dateAdapter={AdapterDateFns}>
          <CssBaseline />
          <BrowserRouter>
            <App />
          </BrowserRouter>
          </LocalizationProvider>
        </ThemeProvider>
      </StyledEngineProvider>
    </QueryClientProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
