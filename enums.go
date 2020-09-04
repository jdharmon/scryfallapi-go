package scryfall

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

        // BorderColors enumerates the values for border colors.
    type BorderColors string

    const (
                // Black ...
        Black BorderColors = "black"
                // Borderless ...
        Borderless BorderColors = "borderless"
                // Gold ...
        Gold BorderColors = "gold"
                // Silver ...
        Silver BorderColors = "silver"
                // White ...
        White BorderColors = "white"
            )
    // PossibleBorderColorsValues returns an array of possible values for the BorderColors const type.
    func PossibleBorderColorsValues() []BorderColors {
        return []BorderColors{Black,Borderless,Gold,Silver,White}
    }

        // Colors enumerates the values for colors.
    type Colors string

    const (
                // B ...
        B Colors = "B"
                // G ...
        G Colors = "G"
                // R ...
        R Colors = "R"
                // U ...
        U Colors = "U"
                // W ...
        W Colors = "W"
            )
    // PossibleColorsValues returns an array of possible values for the Colors const type.
    func PossibleColorsValues() []Colors {
        return []Colors{B,G,R,U,W}
    }

        // Layouts enumerates the values for layouts.
    type Layouts string

    const (
                // Augment ...
        Augment Layouts = "augment"
                // DoubleFacedToken ...
        DoubleFacedToken Layouts = "double_faced_token"
                // Emblem ...
        Emblem Layouts = "emblem"
                // Flip ...
        Flip Layouts = "flip"
                // Host ...
        Host Layouts = "host"
                // Leveler ...
        Leveler Layouts = "leveler"
                // Meld ...
        Meld Layouts = "meld"
                // Normal ...
        Normal Layouts = "normal"
                // Planar ...
        Planar Layouts = "planar"
                // Saga ...
        Saga Layouts = "saga"
                // Scheme ...
        Scheme Layouts = "scheme"
                // Split ...
        Split Layouts = "split"
                // Token ...
        Token Layouts = "token"
                // Transform ...
        Transform Layouts = "transform"
                // Vanguard ...
        Vanguard Layouts = "vanguard"
            )
    // PossibleLayoutsValues returns an array of possible values for the Layouts const type.
    func PossibleLayoutsValues() []Layouts {
        return []Layouts{Augment,DoubleFacedToken,Emblem,Flip,Host,Leveler,Meld,Normal,Planar,Saga,Scheme,Split,Token,Transform,Vanguard}
    }

        // LegalStatus enumerates the values for legal status.
    type LegalStatus string

    const (
                // Banned ...
        Banned LegalStatus = "banned"
                // Legal ...
        Legal LegalStatus = "legal"
                // NotLegal ...
        NotLegal LegalStatus = "not_legal"
                // Restricted ...
        Restricted LegalStatus = "restricted"
            )
    // PossibleLegalStatusValues returns an array of possible values for the LegalStatus const type.
    func PossibleLegalStatusValues() []LegalStatus {
        return []LegalStatus{Banned,Legal,NotLegal,Restricted}
    }

        // Rarity enumerates the values for rarity.
    type Rarity string

    const (
                // Common ...
        Common Rarity = "common"
                // Mythic ...
        Mythic Rarity = "mythic"
                // Rare ...
        Rare Rarity = "rare"
                // Uncommon ...
        Uncommon Rarity = "uncommon"
            )
    // PossibleRarityValues returns an array of possible values for the Rarity const type.
    func PossibleRarityValues() []Rarity {
        return []Rarity{Common,Mythic,Rare,Uncommon}
    }

        // SetTypes enumerates the values for set types.
    type SetTypes string

    const (
                // SetTypesArchenemy ...
        SetTypesArchenemy SetTypes = "archenemy"
                // SetTypesBox ...
        SetTypesBox SetTypes = "box"
                // SetTypesCommander ...
        SetTypesCommander SetTypes = "commander"
                // SetTypesConspiracy ...
        SetTypesConspiracy SetTypes = "conspiracy"
                // SetTypesCore ...
        SetTypesCore SetTypes = "core"
                // SetTypesDuelDeck ...
        SetTypesDuelDeck SetTypes = "duel_deck"
                // SetTypesExpansion ...
        SetTypesExpansion SetTypes = "expansion"
                // SetTypesFromTheVault ...
        SetTypesFromTheVault SetTypes = "from_the_vault"
                // SetTypesFunny ...
        SetTypesFunny SetTypes = "funny"
                // SetTypesMasterpiece ...
        SetTypesMasterpiece SetTypes = "masterpiece"
                // SetTypesMasters ...
        SetTypesMasters SetTypes = "masters"
                // SetTypesMemorabilia ...
        SetTypesMemorabilia SetTypes = "memorabilia"
                // SetTypesPlanechase ...
        SetTypesPlanechase SetTypes = "planechase"
                // SetTypesPremiumDeck ...
        SetTypesPremiumDeck SetTypes = "premium_deck"
                // SetTypesPromo ...
        SetTypesPromo SetTypes = "promo"
                // SetTypesSpellbook ...
        SetTypesSpellbook SetTypes = "spellbook"
                // SetTypesStarter ...
        SetTypesStarter SetTypes = "starter"
                // SetTypesToken ...
        SetTypesToken SetTypes = "token"
                // SetTypesTreasureChest ...
        SetTypesTreasureChest SetTypes = "treasure_chest"
                // SetTypesVanguard ...
        SetTypesVanguard SetTypes = "vanguard"
            )
    // PossibleSetTypesValues returns an array of possible values for the SetTypes const type.
    func PossibleSetTypesValues() []SetTypes {
        return []SetTypes{SetTypesArchenemy,SetTypesBox,SetTypesCommander,SetTypesConspiracy,SetTypesCore,SetTypesDuelDeck,SetTypesExpansion,SetTypesFromTheVault,SetTypesFunny,SetTypesMasterpiece,SetTypesMasters,SetTypesMemorabilia,SetTypesPlanechase,SetTypesPremiumDeck,SetTypesPromo,SetTypesSpellbook,SetTypesStarter,SetTypesToken,SetTypesTreasureChest,SetTypesVanguard}
    }

        // SortDirection enumerates the values for sort direction.
    type SortDirection string

    const (
                // Asc ...
        Asc SortDirection = "asc"
                // Auto ...
        Auto SortDirection = "auto"
                // Desc ...
        Desc SortDirection = "desc"
            )
    // PossibleSortDirectionValues returns an array of possible values for the SortDirection const type.
    func PossibleSortDirectionValues() []SortDirection {
        return []SortDirection{Asc,Auto,Desc}
    }

        // SortOrder enumerates the values for sort order.
    type SortOrder string

    const (
                // SortOrderArtist ...
        SortOrderArtist SortOrder = "artist"
                // SortOrderCmc ...
        SortOrderCmc SortOrder = "cmc"
                // SortOrderColor ...
        SortOrderColor SortOrder = "color"
                // SortOrderEdhrec ...
        SortOrderEdhrec SortOrder = "edhrec"
                // SortOrderEur ...
        SortOrderEur SortOrder = "eur"
                // SortOrderName ...
        SortOrderName SortOrder = "name"
                // SortOrderPower ...
        SortOrderPower SortOrder = "power"
                // SortOrderRarity ...
        SortOrderRarity SortOrder = "rarity"
                // SortOrderReleased ...
        SortOrderReleased SortOrder = "released"
                // SortOrderSet ...
        SortOrderSet SortOrder = "set"
                // SortOrderTix ...
        SortOrderTix SortOrder = "tix"
                // SortOrderToughness ...
        SortOrderToughness SortOrder = "toughness"
                // SortOrderUsd ...
        SortOrderUsd SortOrder = "usd"
            )
    // PossibleSortOrderValues returns an array of possible values for the SortOrder const type.
    func PossibleSortOrderValues() []SortOrder {
        return []SortOrder{SortOrderArtist,SortOrderCmc,SortOrderColor,SortOrderEdhrec,SortOrderEur,SortOrderName,SortOrderPower,SortOrderRarity,SortOrderReleased,SortOrderSet,SortOrderTix,SortOrderToughness,SortOrderUsd}
    }

        // UniqueStrategy enumerates the values for unique strategy.
    type UniqueStrategy string

    const (
                // Art ...
        Art UniqueStrategy = "art"
                // Cards ...
        Cards UniqueStrategy = "cards"
                // Prints ...
        Prints UniqueStrategy = "prints"
            )
    // PossibleUniqueStrategyValues returns an array of possible values for the UniqueStrategy const type.
    func PossibleUniqueStrategyValues() []UniqueStrategy {
        return []UniqueStrategy{Art,Cards,Prints}
    }
