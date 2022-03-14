import { Container } from "@mui/material";
import { Navbar } from "components/organisms/navbar/navbar";
import React from "react";
import { Route, Routes } from "react-router-dom";
import { FeedPage } from "components/pages/feed";
import { ProfilePage } from "components/pages/profile";
import { ResultsPage } from "components/pages/results";
import { FullsizeEventCard } from "components/pages/fullsize-event-card";

const App = (): JSX.Element => {
  return (
    <>
      <Container maxWidth="sm">
        <Routes>
          <Route path="/" element={<FeedPage />} />
          <Route path="/feed" element={<FeedPage />} />
          <Route path="/profile" element={<ProfilePage />}/>
          <Route path="/results" element={<ResultsPage />}/>
          <Route path="/event" element={<FullsizeEventCard />} />
        </Routes>
        <Navbar />
      </Container>
    </>
  );
};

export default App;
