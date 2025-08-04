# Concept


```mermaid
---
title: Play Concept
---
flowchart LR
    A[City]
    B[Policy]
    C[Economy]
    D[People]
    E[Resources]
    F[Event]
    A --> B
    A --> C
    A --> D
    A --> E
    A --> F
    B1[Tax]
    B2[Regulation]
    B3[Technology]
    B --> B1
    B --> B2
    B --> B3
    C1[Trade]
    C2[Demand]
    C3[Loan]
    C --> C1
    C --> C2
    C --> C3
    D1[Citizens]
    D2[Advisors]
    D --> D1
    D --> D2
    E1[Water]
    E2[Electricity]
    E3[Food]
    E4[Goods]
    E5[Cash]
    E --> E1
    E --> E2
    E --> E3
    E --> E4
    E --> E5
    F1[Natural Disaster]
    F2[Upgrade]
    F --> F1
    F --> F2
```

## Policy

The policy section of the game mechanics includes research and push regulation and stuff.

So we will have tech tree and each tech tree can have an interface with

```typescript
interface Technology {
    name: string
    done: boolean
    progress: number
    next: Array<Technology>
    last: Technology | undefined
}

interface TechTree {
    get: (name:string) => Technology | undefined
}
```

This interface can let user know if the tech is done or not\
The technology research will unlock game mechanics as time pass

When the goverment and economy unlock to the certain level, the bank and tax policy will shows up

## Economy

## People

Each person in the game is a individual living being, The major (Player) cannot directly control to it.

People needs food and watch to survive, To the certain satisfaction level, it will go reproduce

## Resources

* Food
* Watch
* Goods
* Electricity
* Cash
* etc...

At the begining, The cash field will have gold bar symbol, This means there is no money yet.

## Event