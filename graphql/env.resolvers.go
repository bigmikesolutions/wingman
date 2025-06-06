package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"

	"github.com/bigmikesolutions/wingman/graphql/conv"
	"github.com/bigmikesolutions/wingman/graphql/generated"
	"github.com/bigmikesolutions/wingman/graphql/model"
)

// EnvGrant is the resolver for the envGrant field.
func (r *mutationResolver) EnvGrant(ctx context.Context, input model.EnvGrantInput) (*model.EnvGrantPayload, error) {
	token, err := r.Tokens.Create(map[string]any{})
	if err != nil {
		r.reqLog(ctx).Error().Err(err).Msg("Env grant error")
		return nil, err
	}

	return &model.EnvGrantPayload{
		MutationID: input.MutationID,
		Token:      &token,
	}, nil
}

// CreateEnvironment is the resolver for the createEnvironment field.
func (r *mutationResolver) CreateEnvironment(ctx context.Context, input model.CreateEnvironmentInput) (*model.CreateEnvironmentPayload, error) {
	e := conv.CreateEnvironmentInputToInternal(input)
	if err := r.Environments.Create(ctx, e); err != nil {
		return nil, err
	}

	return &model.CreateEnvironmentPayload{
		MutationID: input.MutationID,
		Env:        e.ID,
	}, nil
}

// Environment is the resolver for the environment field.
func (r *queryResolver) Environment(ctx context.Context, id string) (*model.Environment, error) {
	e, err := r.Environments.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return conv.EnvInternalToPublic(e), nil
}

// Environment returns generated.EnvironmentResolver implementation.
func (r *Resolver) Environment() generated.EnvironmentResolver { return &environmentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// ResourceGrantInput returns generated.ResourceGrantInputResolver implementation.
func (r *Resolver) ResourceGrantInput() generated.ResourceGrantInputResolver {
	return &resourceGrantInputResolver{r}
}

type (
	environmentResolver        struct{ *Resolver }
	mutationResolver           struct{ *Resolver }
	queryResolver              struct{ *Resolver }
	resourceGrantInputResolver struct{ *Resolver }
)
