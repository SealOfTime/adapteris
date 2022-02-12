import { Container } from "@mui/material";
import { Navbar } from "components/organisms/navbar/navbar";
import React from "react";
import { Route, Routes } from "react-router-dom";
import { FeedPage } from "components/pages/feed";
import { ProfilePage } from "components/pages/profile";

const App = (): JSX.Element => {
  return (
    <>
      <Container maxWidth="sm">
        <Routes>
          <Route path="/feed" element={<FeedPage />} />
          <Route />
          <Route path="/profile" element={<ProfilePage />}/>
        </Routes>
        <Navbar />
      </Container>
    </>
  );
};

export default App;
