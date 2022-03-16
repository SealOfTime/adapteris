import {Box, Container} from "@mui/material";
import {Navbar} from "components/organisms/navbar/navbar";
import React from "react";
import {Route, Routes} from "react-router-dom";
import {FeedPage} from "components/pages/feed";
import {ProfilePage} from "components/pages/profile";
import {ResultsPage} from "components/pages/results";
import {LoginPage} from "components/pages/login";
import {SchoolPage} from "components/pages/school";
import {EventPage} from "components/pages/event-page";
import {LoggedInPage} from "components/pages/loggedin";
import {SchoolPreviewPage} from "components/pages/school-preview";

const App = (): JSX.Element => {
  return (
    <>
      <Container maxWidth="sm" sx={{height: '100vh'}}>
        <Box pb="64px" sx={{height: '100%'}}>
            <Routes>
              <Route path="/" element={<SchoolPreviewPage />} />
              <Route path="/loggedin" element={<LoggedInPage/>}/>
              <Route path="/school">
                <Route index element={<SchoolPage />} />
                <Route path=":schoolId" element={<SchoolPage/>}/>
              </Route>
              <Route path="/feed" element={<FeedPage />} />
              <Route path="/profile">
                <Route index element={<ProfilePage />} />
              </Route>
              <Route path="/results" element={<ResultsPage />} />
              <Route path="/event/:eventId" element={<EventPage />} />
              <Route path="/login" element={<LoginPage />} />
            </Routes>
        </Box>
        <Navbar />
      </Container>
    </>
  );
};

export default App;
