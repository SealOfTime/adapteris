import React, {FC} from "react";
import {Controller} from "react-hook-form";
import {DatePicker} from "@mui/lab";
import {TextField} from "@mui/material";

export const HookFormDatePicker: FC<any> = <T, Ctx>({name, label, control}) => (
    <Controller
        name={name}
        control={control}
        render={({field: {value, onChange}}) =>
            <DatePicker
                label={label}
                value={value}
                onChange={v => onChange(v)}
                renderInput={(params) => <TextField {...params} />}
            />
        }
    />
)
