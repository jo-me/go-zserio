// Code generated from /private/var/tmp/_bazel_ignas.anikevicius/c3129ead79ca509099c649d74caf832e/sandbox/darwin-sandbox/2595/execroot/__main__/internal/parser/ZserioParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ZserioParser
import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by ZserioParser.
type ZserioParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ZserioParser#packageDeclaration.
	VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#packageNameDefinition.
	VisitPackageNameDefinition(ctx *PackageNameDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#languageDirective.
	VisitLanguageDirective(ctx *LanguageDirectiveContext) interface{}

	// Visit a parse tree produced by ZserioParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#symbolDefinition.
	VisitSymbolDefinition(ctx *SymbolDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#constDefinition.
	VisitConstDefinition(ctx *ConstDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#ruleGroupDefinition.
	VisitRuleGroupDefinition(ctx *RuleGroupDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#ruleDefinition.
	VisitRuleDefinition(ctx *RuleDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#subtypeDeclaration.
	VisitSubtypeDeclaration(ctx *SubtypeDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#structureDeclaration.
	VisitStructureDeclaration(ctx *StructureDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#structureFieldDefinition.
	VisitStructureFieldDefinition(ctx *StructureFieldDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#fieldAlignment.
	VisitFieldAlignment(ctx *FieldAlignmentContext) interface{}

	// Visit a parse tree produced by ZserioParser#fieldOffset.
	VisitFieldOffset(ctx *FieldOffsetContext) interface{}

	// Visit a parse tree produced by ZserioParser#fieldTypeId.
	VisitFieldTypeId(ctx *FieldTypeIdContext) interface{}

	// Visit a parse tree produced by ZserioParser#fieldArrayRange.
	VisitFieldArrayRange(ctx *FieldArrayRangeContext) interface{}

	// Visit a parse tree produced by ZserioParser#fieldInitializer.
	VisitFieldInitializer(ctx *FieldInitializerContext) interface{}

	// Visit a parse tree produced by ZserioParser#fieldOptionalClause.
	VisitFieldOptionalClause(ctx *FieldOptionalClauseContext) interface{}

	// Visit a parse tree produced by ZserioParser#fieldConstraint.
	VisitFieldConstraint(ctx *FieldConstraintContext) interface{}

	// Visit a parse tree produced by ZserioParser#choiceDeclaration.
	VisitChoiceDeclaration(ctx *ChoiceDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#choiceCases.
	VisitChoiceCases(ctx *ChoiceCasesContext) interface{}

	// Visit a parse tree produced by ZserioParser#choiceCase.
	VisitChoiceCase(ctx *ChoiceCaseContext) interface{}

	// Visit a parse tree produced by ZserioParser#choiceDefault.
	VisitChoiceDefault(ctx *ChoiceDefaultContext) interface{}

	// Visit a parse tree produced by ZserioParser#choiceFieldDefinition.
	VisitChoiceFieldDefinition(ctx *ChoiceFieldDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#unionDeclaration.
	VisitUnionDeclaration(ctx *UnionDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#unionFieldDefinition.
	VisitUnionFieldDefinition(ctx *UnionFieldDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#enumItem.
	VisitEnumItem(ctx *EnumItemContext) interface{}

	// Visit a parse tree produced by ZserioParser#bitmaskDeclaration.
	VisitBitmaskDeclaration(ctx *BitmaskDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#bitmaskValue.
	VisitBitmaskValue(ctx *BitmaskValueContext) interface{}

	// Visit a parse tree produced by ZserioParser#sqlTableDeclaration.
	VisitSqlTableDeclaration(ctx *SqlTableDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#sqlTableFieldDefinition.
	VisitSqlTableFieldDefinition(ctx *SqlTableFieldDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#sqlConstraintDefinition.
	VisitSqlConstraintDefinition(ctx *SqlConstraintDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#sqlConstraint.
	VisitSqlConstraint(ctx *SqlConstraintContext) interface{}

	// Visit a parse tree produced by ZserioParser#sqlWithoutRowId.
	VisitSqlWithoutRowId(ctx *SqlWithoutRowIdContext) interface{}

	// Visit a parse tree produced by ZserioParser#sqlDatabaseDefinition.
	VisitSqlDatabaseDefinition(ctx *SqlDatabaseDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#sqlDatabaseFieldDefinition.
	VisitSqlDatabaseFieldDefinition(ctx *SqlDatabaseFieldDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#serviceDefinition.
	VisitServiceDefinition(ctx *ServiceDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#serviceMethodDefinition.
	VisitServiceMethodDefinition(ctx *ServiceMethodDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#pubsubDefinition.
	VisitPubsubDefinition(ctx *PubsubDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#pubsubMessageDefinition.
	VisitPubsubMessageDefinition(ctx *PubsubMessageDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#topicDefinition.
	VisitTopicDefinition(ctx *TopicDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by ZserioParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by ZserioParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by ZserioParser#typeParameters.
	VisitTypeParameters(ctx *TypeParametersContext) interface{}

	// Visit a parse tree produced by ZserioParser#parameterDefinition.
	VisitParameterDefinition(ctx *ParameterDefinitionContext) interface{}

	// Visit a parse tree produced by ZserioParser#templateParameters.
	VisitTemplateParameters(ctx *TemplateParametersContext) interface{}

	// Visit a parse tree produced by ZserioParser#templateArguments.
	VisitTemplateArguments(ctx *TemplateArgumentsContext) interface{}

	// Visit a parse tree produced by ZserioParser#templateArgument.
	VisitTemplateArgument(ctx *TemplateArgumentContext) interface{}

	// Visit a parse tree produced by ZserioParser#instantiateDeclaration.
	VisitInstantiateDeclaration(ctx *InstantiateDeclarationContext) interface{}

	// Visit a parse tree produced by ZserioParser#bitwiseXorExpression.
	VisitBitwiseXorExpression(ctx *BitwiseXorExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#dotExpression.
	VisitDotExpression(ctx *DotExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#valueofExpression.
	VisitValueofExpression(ctx *ValueofExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#arrayExpression.
	VisitArrayExpression(ctx *ArrayExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#numbitsExpression.
	VisitNumbitsExpression(ctx *NumbitsExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#lengthofExpression.
	VisitLengthofExpression(ctx *LengthofExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#identifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#bitwiseOrExpression.
	VisitBitwiseOrExpression(ctx *BitwiseOrExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#parenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#bitwiseAndExpression.
	VisitBitwiseAndExpression(ctx *BitwiseAndExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#indexExpression.
	VisitIndexExpression(ctx *IndexExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#literalExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#ternaryExpression.
	VisitTernaryExpression(ctx *TernaryExpressionContext) interface{}

	// Visit a parse tree produced by ZserioParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ZserioParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by ZserioParser#typeReference.
	VisitTypeReference(ctx *TypeReferenceContext) interface{}

	// Visit a parse tree produced by ZserioParser#typeInstantiation.
	VisitTypeInstantiation(ctx *TypeInstantiationContext) interface{}

	// Visit a parse tree produced by ZserioParser#builtinType.
	VisitBuiltinType(ctx *BuiltinTypeContext) interface{}

	// Visit a parse tree produced by ZserioParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by ZserioParser#typeArguments.
	VisitTypeArguments(ctx *TypeArgumentsContext) interface{}

	// Visit a parse tree produced by ZserioParser#typeArgument.
	VisitTypeArgument(ctx *TypeArgumentContext) interface{}

	// Visit a parse tree produced by ZserioParser#dynamicLengthArgument.
	VisitDynamicLengthArgument(ctx *DynamicLengthArgumentContext) interface{}

	// Visit a parse tree produced by ZserioParser#intType.
	VisitIntType(ctx *IntTypeContext) interface{}

	// Visit a parse tree produced by ZserioParser#varintType.
	VisitVarintType(ctx *VarintTypeContext) interface{}

	// Visit a parse tree produced by ZserioParser#fixedBitFieldType.
	VisitFixedBitFieldType(ctx *FixedBitFieldTypeContext) interface{}

	// Visit a parse tree produced by ZserioParser#dynamicBitFieldType.
	VisitDynamicBitFieldType(ctx *DynamicBitFieldTypeContext) interface{}

	// Visit a parse tree produced by ZserioParser#boolType.
	VisitBoolType(ctx *BoolTypeContext) interface{}

	// Visit a parse tree produced by ZserioParser#stringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by ZserioParser#floatType.
	VisitFloatType(ctx *FloatTypeContext) interface{}

	// Visit a parse tree produced by ZserioParser#externType.
	VisitExternType(ctx *ExternTypeContext) interface{}

}