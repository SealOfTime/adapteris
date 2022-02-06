import { Container } from "@mui/material";
import { Navbar } from "components/organisms/navbar/navbar";
import React from "react";
import { Route, Routes } from "react-router-dom";
import { FeedPage } from "components/pages/feed";

const App = (): JSX.Element => {
  return (
    <>
      <Navbar />
      <Container maxWidth="sm">
        <Routes>
          <Route path="/feed" element={<FeedPage />} />
          <Route />
          {/* <Route path="*" element={<NotFound />} /> */}
        </Routes>
      </Container>
    </>
  );
};

export default App;
