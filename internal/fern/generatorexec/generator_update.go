// Generated by Fern. Do not edit.

package generatorexec

import (
	json "encoding/json"
	fmt "fmt"
)

type GeneratorUpdate struct {
	Type             string
	Init             *InitUpdate
	InitV2           *InitUpdateV2
	Log              *LogUpdate
	Publishing       *PackageCoordinate
	Published        *PackageCoordinate
	ExitStatusUpdate *ExitStatusUpdate
}

func NewGeneratorUpdateFromInit(value *InitUpdate) *GeneratorUpdate {
	return &GeneratorUpdate{Type: "init", Init: value}
}

func NewGeneratorUpdateFromInitV2(value *InitUpdateV2) *GeneratorUpdate {
	return &GeneratorUpdate{Type: "initV2", InitV2: value}
}

func NewGeneratorUpdateFromLog(value *LogUpdate) *GeneratorUpdate {
	return &GeneratorUpdate{Type: "log", Log: value}
}

func NewGeneratorUpdateFromPublishing(value *PackageCoordinate) *GeneratorUpdate {
	return &GeneratorUpdate{Type: "publishing", Publishing: value}
}

func NewGeneratorUpdateFromPublished(value *PackageCoordinate) *GeneratorUpdate {
	return &GeneratorUpdate{Type: "published", Published: value}
}

func NewGeneratorUpdateFromExitStatusUpdate(value *ExitStatusUpdate) *GeneratorUpdate {
	return &GeneratorUpdate{Type: "exitStatusUpdate", ExitStatusUpdate: value}
}

func (g *GeneratorUpdate) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	g.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "init":
		value := new(InitUpdate)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Init = value
	case "initV2":
		value := new(InitUpdateV2)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.InitV2 = value
	case "log":
		value := new(LogUpdate)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Log = value
	case "publishing":
		var valueUnmarshaler struct {
			Publishing *PackageCoordinate `json:"publishing"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		g.Publishing = valueUnmarshaler.Publishing
	case "published":
		var valueUnmarshaler struct {
			Published *PackageCoordinate `json:"published"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		g.Published = valueUnmarshaler.Published
	case "exitStatusUpdate":
		var valueUnmarshaler struct {
			ExitStatusUpdate *ExitStatusUpdate `json:"exitStatusUpdate"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		g.ExitStatusUpdate = valueUnmarshaler.ExitStatusUpdate
	}
	return nil
}

func (g GeneratorUpdate) MarshalJSON() ([]byte, error) {
	switch g.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", g.Type, g)
	case "init":
		var marshaler = struct {
			Type string `json:"_type"`
			*InitUpdate
		}{
			Type:       g.Type,
			InitUpdate: g.Init,
		}
		return json.Marshal(marshaler)
	case "initV2":
		var marshaler = struct {
			Type string `json:"_type"`
			*InitUpdateV2
		}{
			Type:         g.Type,
			InitUpdateV2: g.InitV2,
		}
		return json.Marshal(marshaler)
	case "log":
		var marshaler = struct {
			Type string `json:"_type"`
			*LogUpdate
		}{
			Type:      g.Type,
			LogUpdate: g.Log,
		}
		return json.Marshal(marshaler)
	case "publishing":
		var marshaler = struct {
			Type       string             `json:"_type"`
			Publishing *PackageCoordinate `json:"publishing"`
		}{
			Type:       g.Type,
			Publishing: g.Publishing,
		}
		return json.Marshal(marshaler)
	case "published":
		var marshaler = struct {
			Type      string             `json:"_type"`
			Published *PackageCoordinate `json:"published"`
		}{
			Type:      g.Type,
			Published: g.Published,
		}
		return json.Marshal(marshaler)
	case "exitStatusUpdate":
		var marshaler = struct {
			Type             string            `json:"_type"`
			ExitStatusUpdate *ExitStatusUpdate `json:"exitStatusUpdate"`
		}{
			Type:             g.Type,
			ExitStatusUpdate: g.ExitStatusUpdate,
		}
		return json.Marshal(marshaler)
	}
}

type GeneratorUpdateVisitor interface {
	VisitInit(*InitUpdate) error
	VisitInitV2(*InitUpdateV2) error
	VisitLog(*LogUpdate) error
	VisitPublishing(*PackageCoordinate) error
	VisitPublished(*PackageCoordinate) error
	VisitExitStatusUpdate(*ExitStatusUpdate) error
}

func (g *GeneratorUpdate) Accept(v GeneratorUpdateVisitor) error {
	switch g.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", g.Type, g)
	case "init":
		return v.VisitInit(g.Init)
	case "initV2":
		return v.VisitInitV2(g.InitV2)
	case "log":
		return v.VisitLog(g.Log)
	case "publishing":
		return v.VisitPublishing(g.Publishing)
	case "published":
		return v.VisitPublished(g.Published)
	case "exitStatusUpdate":
		return v.VisitExitStatusUpdate(g.ExitStatusUpdate)
	}
}
