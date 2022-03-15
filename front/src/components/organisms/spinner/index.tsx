import React from 'react';

export const Spinner = () => (
    <svg style={{
        margin: 'auto', background: 'none', display: 'block', shapeRendering: 'auto',
    }} width="200px" height="200px"
        viewBox="0 0 100 100" preserveAspectRatio="xMidYMid">
        <g transform="rotate(0 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#93dbe9">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.8899676375404532s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(30 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#689cc5">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.8090614886731392s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(60 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#5e6fa3">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.7281553398058253s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(90 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#3b4368">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.6472491909385114s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(120 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#191d3a">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.5663430420711975s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(150 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#d9dbee">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.4854368932038835s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(180 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#b3b7e2">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.4045307443365696s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(210 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#93dbe9">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.3236245954692557s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(240 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#689cc5">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.24271844660194175s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(270 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#5e6fa3">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.16181229773462785s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(300 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#3b4368">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s"
                    begin="-0.08090614886731393s" repeatCount="indefinite"></animate>
            </rect>
        </g>
        <g transform="rotate(330 50 50)">
            <rect x="47" y="18" rx="3" ry="6" width="6" height="12" fill="#191d3a">
                <animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="0.970873786407767s" begin="0s"
                    repeatCount="indefinite"></animate>
            </rect>
        </g>
    </svg>
)