package sample

import "github.com/ggeorgiev/neo/src/db"

var (
	Topics = []db.Topic{
		{
			ID:   "Vegetarianism",
			Text: "Humans that do not it meat",
		},
	}

	Statements = []db.Statement{
		{
			ID:   "EveryoneVegetarian",
			Text: "All humans should be vegetarian.",
		},
		{
			ID:   "GovermentVegetarian",
			Text: "Goverment should create instentives for people to become vegiterians.",
		},
		{
			ID:   "VegetarianIsChoice",
			Text: "It is everyones choice to be a vegiterian.",
		},
		{
			ID:   "GovermentShouldSubsidies",
			Text: "Goverment should subsidies for animal agriculture.",
		},
		{
			ID:   "NobodyVegetarian",
			Text: "No one should be vegetarian.",
		},
		{
			ID:   "EthicalTreatment",
			Text: "It is ethically wrong to harm animals.",
		},
		{
			ID:   "FactoryFarming",
			Text: "Factory farming practices are cruel and cause undue suffering to animals.",
		},
		{
			ID:   "PlantNutrition",
			Text: "Humans can obtain all necessary nutrients from plant-based sources.",
		},
		{
			ID:   "OmnivorousEvolution",
			Text: "Humans have evolved as omnivores, and it's natural to consume both plant and animal products.",
		},
		{
			ID:   "HistoricalMeatConsumption",
			Text: "Historical evidence shows that early humans consumed meat as part of their diet.",
		},
		{
			ID:   "B12Source",
			Text: "Some nutrients like B12 are primarily found in animal products and are essential for human health.",
		},
		{
			ID:   "AnimalSentience",
			Text: "Animals are sentient beings capable of experiencing pain and emotions.",
		},
		{
			ID:   "UnnaturalAnimalTreatment",
			Text: "Many animals in the meat industry are subjected to unnatural and harsh treatments.",
		},
		{
			ID:   "NoMoralDifference",
			Text: "There's no moral difference between causing harm to a human and causing harm to an animal.",
		},
	}

	Arguments = []db.Argument{
		{
			ID:   "EthicalTreatmentOfAnimals",
			Text: "It is ethically wrong to harm animals.",
		},
		{
			ID:   "HumansEvolvedAsOmnivores",
			Text: "Humans have evolved as omnivores, and it's natural to consume both plant and animal products.",
		},
		{
			ID:   "HealthBenefits",
			Text: "A vegetarian diet leads to numerous health benefits including lower risk of chronic diseases.",
		},
		{
			ID:   "EnvironmentalConcerns",
			Text: "Vegetarian diets have a lower environmental impact than meat-based diets.",
		},
		{
			ID:   "EconomicBenefits",
			Text: "Raising livestock is resource-intensive and not economically sustainable.",
		},
		{
			ID:   "AnimalRights",
			Text: "Animals have a right to life and should not be raised for slaughter.",
		},
		{
			ID:   "ReligiousBeliefs",
			Text: "Many religions promote vegetarianism or discourage harm to animals.",
		},
		{
			ID:   "NaturalFoodChain",
			Text: "Animals eat other animals; it's the natural food chain.",
		},
		{
			ID:   "PersonalChoice",
			Text: "Dietary choices are a personal decision and should not be imposed on others.",
		},
		{
			ID:   "CulturalTraditions",
			Text: "Meat consumption is deeply rooted in many cultures' traditions and festivities.",
		},
		{
			ID:   "BiologicalNeeds",
			Text: "Humans have specific biological needs that may require nutrients only found in animal products.",
		},
		{
			ID:   "EconomicDisruption",
			Text: "A sudden shift to vegetarianism can disrupt the economy, especially in regions reliant on the meat industry.",
		},
		{
			ID:   "PlantBasedEarlyHumans",
			Text: "Early humans primarily consumed a plant-based diet before the advent of hunting.",
		},
		{
			ID:   "ModernHealthIssues",
			Text: "Modern health issues arise from the excessive consumption of meat and animal products.",
		},
		{
			ID:   "TeethEvolution",
			Text: "The human dental structure is suited for both plant and meat consumption.",
		},
	}

	Partains = []db.Partain{
		{
			TopicID:     "Vegetarianism",
			StatementID: "EveryoneVegetarian",
		},
		{
			TopicID:     "Vegetarianism",
			StatementID: "VegetarianIsChoice",
		},
		{
			TopicID:     "Vegetarianism",
			StatementID: "NobodyVegetarian",
		},
	}

	Stances = []db.Stance{
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "EthicalTreatmentOfAnimals",
			Type:        db.For,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "HumansEvolvedAsOmnivores",
			Type:        db.Against,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "HealthBenefits",
			Type:        db.For,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "EnvironmentalConcerns",
			Type:        db.For,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "EconomicBenefits",
			Type:        db.For,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "AnimalRights",
			Type:        db.For,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "ReligiousBeliefs",
			Type:        db.For,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "NaturalFoodChain",
			Type:        db.Against,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "PersonalChoice",
			Type:        db.Against,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "CulturalTraditions",
			Type:        db.Against,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "BiologicalNeeds",
			Type:        db.Against,
		},
		{
			StatementID: "EveryoneVegetarian",
			ArgumentID:  "EconomicDisruption",
			Type:        db.Against,
		},
	}

	Evidences = []db.Evidence{
		{
			StatementID: "FactoryFarming",
			ArgumentID:  "EthicalTreatmentOfAnimals",
			Type:        db.Supports,
		},
		{
			StatementID: "HistoricalMeatConsumption",
			ArgumentID:  "HumansEvolvedAsOmnivores",
			Type:        db.Contradicts,
		},
		{
			StatementID: "AnimalSentience",
			ArgumentID:  "EthicalTreatmentOfAnimals",
			Type:        db.Supports,
		},
		{
			StatementID: "UnnaturalAnimalTreatment",
			ArgumentID:  "EthicalTreatmentOfAnimals",
			Type:        db.Supports,
		},
		{
			StatementID: "NoMoralDifference",
			ArgumentID:  "EthicalTreatmentOfAnimals",
			Type:        db.Supports,
		},
		{
			StatementID: "PlantBasedEarlyHumans",
			ArgumentID:  "HumansEvolvedAsOmnivores",
			Type:        db.Contradicts,
		},
		{
			StatementID: "ModernHealthIssues",
			ArgumentID:  "HumansEvolvedAsOmnivores",
			Type:        db.Contradicts,
		},
		{
			StatementID: "TeethEvolution",
			ArgumentID:  "HumansEvolvedAsOmnivores",
			Type:        db.Supports,
		},
	}
)
