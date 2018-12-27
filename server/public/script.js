(function(document, data) {
    "use strict";

    const initialColors = [
        [22, 91, 136],
        [240, 65, 44],
        [3, 165, 150],
        [255, 128, 0],
        [240, 170, 33]
    ];

    const options = {
        legend: { display: false },
        title: { display: false }
    };

    function colors(participants = []) {
        if (participants.length < initialColors.length) {
            return participants
                .map((p, i) => ({ [p]: initialColors[i] }))
                .reduce((p, n) => ({ ...p, ...n }), {});
        }

        return participants
            .map((p, i) => ({
                [p]:
                    typeof initialColors[i] !== "undefined"
                        ? initialColors[i]
                        : [
                            Math.floor(Math.random() * 256),
                            Math.floor(Math.random() * 256),
                            Math.floor(Math.random() * 256)
                        ]
            }))
            .reduce((p, n) => ({ ...p, ...n }), {});
    }

    const colorsMap = colors(data.participants);

    function rgbaFormat([r, g, b], a = 1) {
        return `rgba(${r},${g},${b},${a})`;
    }

    function generateLegends() {
        const html = Object.entries(colorsMap)
            .map(
                ([name, color]) =>
                    `<div class="legend"><span style="background: ${rgbaFormat(
                        color
                    )};"></span><p>${name}</p></div>`
            )
            .join("");

        return (document
            .getElementsByClassName("legends")
            .item(0).innerHTML = html);
    }

    function generateFrequency(id = "radar") {
        const labels = Object.keys(data.frequency);
        const common = { lineTension: 0.1, borderWidth: 2 };
        const dataset = Object.entries(data.contributionFrequency).map(
            ([key, value]) => ({
                ...common,
                backgroundColor: rgbaFormat(colorsMap[key], 0.5),
                borderColor: rgbaFormat(colorsMap[key]),
                label: key,
                data: Object.entries(value).map(([, v]) => v)
            })
        );

        return new Chart(id, {
            type: "radar",
            data: {
                labels,
                datasets: [
                    ...dataset,
                    {
                        ...common,
                        data: Object.entries(data.frequency).map(([, value]) => value),
                        label: "Mutual"
                    }
                ]
            },
            options: {
                ...options
            }
        });
    }

    function generateContributions(prefix = "contrib") {
        return [
            {
                suffix: "count",
                src: Object.entries(data.contributionCount)
            },
            {
                suffix: "words",
                src: Object.entries(data.contributionWords)
            },
            {
                suffix: "letters",
                src: Object.entries(data.contributionLetters)
            }
        ].map(
            model =>
                new Chart(`${prefix}-${model.suffix}`, {
                    type: "polarArea",
                    data: {
                        labels: model.src.map(([key]) => key),
                        datasets: [
                            {
                                backgroundColor:model.src.map(([key]) => rgbaFormat(colorsMap[key])),
                                data: model.src.map(([, value]) => value)
                            }
                        ]
                    },
                    options: {
                        ...options,
                        maintainAspectRatio: false,
                        responsive: true
                    }
                })
        );
    }

    function generateTimeline(id = "timeline") {
        const keys = Object.keys(data.timeline);
        const labels = keys.map(k => k.split("T")[0]);
        const raw = keys.map(k => ({
            ...data.participants
                .map(p => ({
                    [p]: data.timeline[k].filter(t => t.sender === p).length
                }))
                .reduce((p, n) => ({ ...p, ...n }), {})
        }));

        document
            .getElementsByClassName("timeline")
            .item(0)
            .setAttribute("style", `height: ${labels.length * 3}vh`);
        return new Chart(id, {
            type: "horizontalBar",
            data: {
                labels,
                datasets: data.participants.map(p => ({
                    backgroundColor: rgbaFormat(colorsMap[p]),
                    label: p,
                    data: raw.map(r => r[p])
                }))
            },
            options: {
                ...options,
                maintainAspectRatio: false,
                responsive: true,
                scales: {
                    xAxes: [
                        {
                            stacked: true
                        }
                    ],
                    yAxes: [
                        {
                            stacked: true
                        }
                    ]
                }
            }
        });
    }

    function generateEmojis(id = "emojis") {
        const raw = Array.from(
            new Set(
                Object.entries(data.emojisUsed)
                    .map(([k, v]) => Object.keys(v))
                    .reduce((p, n) => [...p, ...n], [])
            )
        )
            .map(l => ({
                [l]: data.participants
                    .map(p => ({
                        [p]:
                            typeof data.emojisUsed[p][l] !== "undefined"
                                ? data.emojisUsed[p][l]
                                : 0
                    }))
                    .reduce((p, n) => ({ ...p, ...n }), {})
            }))
            .reduce((p, n) => ({ ...p, ...n }), {});
        const labels = Object.keys(raw);

        document
            .getElementsByClassName("emojis")
            .item(0)
            .setAttribute("style", `height: ${labels.length * 3}vh`);

        return new Chart(id, {
            type: "horizontalBar",
            data: {
                labels,
                datasets: data.participants.map(p => ({
                    label: p,
                    backgroundColor: rgbaFormat(colorsMap[p]),
                    data: labels.map(l => raw[l][p])
                }))
            },
            options: {
                ...options,
                maintainAspectRatio: false,
                responsive: true,
                scales: {
                    xAxes: [
                        {
                            stacked: true
                        }
                    ],
                    yAxes: [
                        {
                            stacked: true,
                            ticks: {
                                fontSize: 22
                            }
                        }
                    ]
                }
            }
        });
    }

    generateLegends();
    generateFrequency();
    generateTimeline();
    generateContributions();
    generateEmojis();
})(window.document, window.data);
