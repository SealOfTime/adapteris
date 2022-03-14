import { Typography, Box, Stack, Card } from "@mui/material"
import { UserEventParticipation } from "components/organisms/result-card/result-card"
import { FC } from "react"
import { EmojiProvider, Emoji } from 'react-apple-emojis'
import emojiData from 'react-apple-emojis/lib/data.json'
import { ResultCard } from "./../organisms/result-card/result-card"

export const ResultsPage: FC = () => {
    const emojiStyle = { width: '3rem' };
    const events = stubUserEventParticipation;
    return (
        <>
            <EmojiProvider data={emojiData}>
                <Box pt='1.25rem'>
                    <Stack width="100%" pt='1.25rem' paddingLeft='0.5rem' direction="row" justifyContent="left" spacing="0.5rem">
                        <Typography variant="h1" gutterBottom>Мои результаты</Typography>
                        <Emoji name="sparkles" style={emojiStyle}></Emoji>
                    </Stack>
                </Box>
                <Stack pb="64px" pt="10px">
                    <Card>
                        <Stack spacing={2} direction="column" justifyContent="space-between">
                            {events.map(e => <ResultCard key={e.name} event={e} />)}
                           </Stack>
                    </Card>
                </Stack>
            </EmojiProvider>
        </>
    );
}

const stubUserEventParticipation: UserEventParticipation[] = [
    {
        name: "Игротехника",
        results: 60
    },
    {
        name: "Публичные выступления",
        results: 24
    },
    {
        name: "Тайм-менеджмент",
        results: 71
    },
    {
        name: "Игротехника",
        results: 100
    },
];