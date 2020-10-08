package gateway

import (
	"context"

	gql "github.com/machinebox/graphql"

	"github.com/railwayapp/cli/entity"
)

func (g *Gateway) GetEnvs(ctx context.Context, req *entity.GetEnvsRequest) (*entity.Envs, error) {
	gqlReq := gql.NewRequest(`
		query ($projectId: String!, $environmentId: String!) {
			allEnvsForEnvironment(projectId: $projectId, environmentId: $environmentId)
		}
	`)
	gqlReq.Var("projectId", req.ProjectID)
	gqlReq.Var("environmentId", req.EnvironmentID)

	err := g.authorize(ctx, gqlReq)

	if err != nil {
		return nil, err
	}

	var resp struct {
		Envs *entity.Envs `json:"allEnvsForEnvironment"`
	}
	if err := g.gqlClient.Run(ctx, gqlReq, &resp); err != nil {
		return nil, err
	}
	return resp.Envs, nil
}

func (g *Gateway) GetEnvsWithEnvToken(ctx context.Context) (*entity.Envs, error) {
	gqlReq := gql.NewRequest(`
	  query {
			allEnvsForEnvToken
		}
	`)

	err := g.setEnvToken(ctx, gqlReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Envs *entity.Envs `json:"allEnvsForEnvToken"`
	}
	if err := g.gqlClient.Run(ctx, gqlReq, &resp); err != nil {
		return nil, err
	}

	return resp.Envs, nil
}
