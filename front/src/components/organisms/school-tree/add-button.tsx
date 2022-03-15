import {IconButton, IconButtonProps, Tooltip} from "@mui/material";
import {PlusCircleIcon} from "@heroicons/react/outline";
import React, {FC} from "react";

export const AddButton: FC<IconButtonProps> = React.forwardRef(
    (props, ref) => (
        <IconButton
            {...props}
            ref={ref}
            sx={{
                width: 64,
                height: 64,
                color: 'grey.400',
                '&:hover': {
                    color: 'grey.500',
                }
            }}
        >
            <PlusCircleIcon />
        </IconButton>
    )
)
