import { Box, Container } from "@mui/material";
import { Navbar } from "components/organisms/navbar/navbar";
import React from "react";
import { Route, Routes } from "react-router-dom";
import { FeedPage } from "components/pages/feed";
import { ProfilePage } from "components/pages/profile";
import { ResultsPage } from "components/pages/results";
import { LoginPage } from "components/pages/login";
import { ProvideAuth } from "hooks/useAuth";
import { SchoolPage } from "components/pages/school";
import { FullsizeEventCard } from "components/pages/fullsize-event-card";

const App = (): JSX.Element => {
  return (
    <>

      <Container maxWidth="sm">
        <Box pb="64px">
          <ProvideAuth>
            <Routes>
              <Route path="/" element={<FeedPage />} />
              <Route path="/school">
                <Route index element={<SchoolPage />} />
                <Route path=":schoolId" element={<SchoolPage/>}/>
              </Route>
              <Route path="/feed" element={<FeedPage />} />
              <Route path="/profile">
                <Route index element={<ProfilePage />} />
              </Route>
              <Route path="/results" element={<ResultsPage />} />
              <Route path="/event/:eventId" element={<FullsizeEventCard />} />
              <Route path="/login" element={<LoginPage />} />
            </Routes>
          </ProvideAuth>
        </Box>
        <Navbar />
      </Container>
    </>
  );
};

export default App;
