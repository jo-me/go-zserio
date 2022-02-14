// Code generated from /private/var/tmp/_bazel_ignas.anikevicius/c3129ead79ca509099c649d74caf832e/sandbox/darwin-sandbox/2595/execroot/__main__/internal/parser/ZserioParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ZserioParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ZserioParserListener is a complete listener for a parse tree produced by ZserioParser.
type ZserioParserListener interface {
	antlr.ParseTreeListener

	// EnterPackageDeclaration is called when entering the packageDeclaration production.
	EnterPackageDeclaration(c *PackageDeclarationContext)

	// EnterPackageNameDefinition is called when entering the packageNameDefinition production.
	EnterPackageNameDefinition(c *PackageNameDefinitionContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterLanguageDirective is called when entering the languageDirective production.
	EnterLanguageDirective(c *LanguageDirectiveContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterSymbolDefinition is called when entering the symbolDefinition production.
	EnterSymbolDefinition(c *SymbolDefinitionContext)

	// EnterConstDefinition is called when entering the constDefinition production.
	EnterConstDefinition(c *ConstDefinitionContext)

	// EnterRuleGroupDefinition is called when entering the ruleGroupDefinition production.
	EnterRuleGroupDefinition(c *RuleGroupDefinitionContext)

	// EnterRuleDefinition is called when entering the ruleDefinition production.
	EnterRuleDefinition(c *RuleDefinitionContext)

	// EnterSubtypeDeclaration is called when entering the subtypeDeclaration production.
	EnterSubtypeDeclaration(c *SubtypeDeclarationContext)

	// EnterStructureDeclaration is called when entering the structureDeclaration production.
	EnterStructureDeclaration(c *StructureDeclarationContext)

	// EnterStructureFieldDefinition is called when entering the structureFieldDefinition production.
	EnterStructureFieldDefinition(c *StructureFieldDefinitionContext)

	// EnterFieldAlignment is called when entering the fieldAlignment production.
	EnterFieldAlignment(c *FieldAlignmentContext)

	// EnterFieldOffset is called when entering the fieldOffset production.
	EnterFieldOffset(c *FieldOffsetContext)

	// EnterFieldTypeId is called when entering the fieldTypeId production.
	EnterFieldTypeId(c *FieldTypeIdContext)

	// EnterFieldArrayRange is called when entering the fieldArrayRange production.
	EnterFieldArrayRange(c *FieldArrayRangeContext)

	// EnterFieldInitializer is called when entering the fieldInitializer production.
	EnterFieldInitializer(c *FieldInitializerContext)

	// EnterFieldOptionalClause is called when entering the fieldOptionalClause production.
	EnterFieldOptionalClause(c *FieldOptionalClauseContext)

	// EnterFieldConstraint is called when entering the fieldConstraint production.
	EnterFieldConstraint(c *FieldConstraintContext)

	// EnterChoiceDeclaration is called when entering the choiceDeclaration production.
	EnterChoiceDeclaration(c *ChoiceDeclarationContext)

	// EnterChoiceCases is called when entering the choiceCases production.
	EnterChoiceCases(c *ChoiceCasesContext)

	// EnterChoiceCase is called when entering the choiceCase production.
	EnterChoiceCase(c *ChoiceCaseContext)

	// EnterChoiceDefault is called when entering the choiceDefault production.
	EnterChoiceDefault(c *ChoiceDefaultContext)

	// EnterChoiceFieldDefinition is called when entering the choiceFieldDefinition production.
	EnterChoiceFieldDefinition(c *ChoiceFieldDefinitionContext)

	// EnterUnionDeclaration is called when entering the unionDeclaration production.
	EnterUnionDeclaration(c *UnionDeclarationContext)

	// EnterUnionFieldDefinition is called when entering the unionFieldDefinition production.
	EnterUnionFieldDefinition(c *UnionFieldDefinitionContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterEnumItem is called when entering the enumItem production.
	EnterEnumItem(c *EnumItemContext)

	// EnterBitmaskDeclaration is called when entering the bitmaskDeclaration production.
	EnterBitmaskDeclaration(c *BitmaskDeclarationContext)

	// EnterBitmaskValue is called when entering the bitmaskValue production.
	EnterBitmaskValue(c *BitmaskValueContext)

	// EnterSqlTableDeclaration is called when entering the sqlTableDeclaration production.
	EnterSqlTableDeclaration(c *SqlTableDeclarationContext)

	// EnterSqlTableFieldDefinition is called when entering the sqlTableFieldDefinition production.
	EnterSqlTableFieldDefinition(c *SqlTableFieldDefinitionContext)

	// EnterSqlConstraintDefinition is called when entering the sqlConstraintDefinition production.
	EnterSqlConstraintDefinition(c *SqlConstraintDefinitionContext)

	// EnterSqlConstraint is called when entering the sqlConstraint production.
	EnterSqlConstraint(c *SqlConstraintContext)

	// EnterSqlWithoutRowId is called when entering the sqlWithoutRowId production.
	EnterSqlWithoutRowId(c *SqlWithoutRowIdContext)

	// EnterSqlDatabaseDefinition is called when entering the sqlDatabaseDefinition production.
	EnterSqlDatabaseDefinition(c *SqlDatabaseDefinitionContext)

	// EnterSqlDatabaseFieldDefinition is called when entering the sqlDatabaseFieldDefinition production.
	EnterSqlDatabaseFieldDefinition(c *SqlDatabaseFieldDefinitionContext)

	// EnterServiceDefinition is called when entering the serviceDefinition production.
	EnterServiceDefinition(c *ServiceDefinitionContext)

	// EnterServiceMethodDefinition is called when entering the serviceMethodDefinition production.
	EnterServiceMethodDefinition(c *ServiceMethodDefinitionContext)

	// EnterPubsubDefinition is called when entering the pubsubDefinition production.
	EnterPubsubDefinition(c *PubsubDefinitionContext)

	// EnterPubsubMessageDefinition is called when entering the pubsubMessageDefinition production.
	EnterPubsubMessageDefinition(c *PubsubMessageDefinitionContext)

	// EnterTopicDefinition is called when entering the topicDefinition production.
	EnterTopicDefinition(c *TopicDefinitionContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterParameterDefinition is called when entering the parameterDefinition production.
	EnterParameterDefinition(c *ParameterDefinitionContext)

	// EnterTemplateParameters is called when entering the templateParameters production.
	EnterTemplateParameters(c *TemplateParametersContext)

	// EnterTemplateArguments is called when entering the templateArguments production.
	EnterTemplateArguments(c *TemplateArgumentsContext)

	// EnterTemplateArgument is called when entering the templateArgument production.
	EnterTemplateArgument(c *TemplateArgumentContext)

	// EnterInstantiateDeclaration is called when entering the instantiateDeclaration production.
	EnterInstantiateDeclaration(c *InstantiateDeclarationContext)

	// EnterBitwiseXorExpression is called when entering the bitwiseXorExpression production.
	EnterBitwiseXorExpression(c *BitwiseXorExpressionContext)

	// EnterDotExpression is called when entering the dotExpression production.
	EnterDotExpression(c *DotExpressionContext)

	// EnterValueofExpression is called when entering the valueofExpression production.
	EnterValueofExpression(c *ValueofExpressionContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterArrayExpression is called when entering the arrayExpression production.
	EnterArrayExpression(c *ArrayExpressionContext)

	// EnterNumbitsExpression is called when entering the numbitsExpression production.
	EnterNumbitsExpression(c *NumbitsExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterLengthofExpression is called when entering the lengthofExpression production.
	EnterLengthofExpression(c *LengthofExpressionContext)

	// EnterIdentifierExpression is called when entering the identifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterBitwiseOrExpression is called when entering the bitwiseOrExpression production.
	EnterBitwiseOrExpression(c *BitwiseOrExpressionContext)

	// EnterParenthesizedExpression is called when entering the parenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterBitwiseAndExpression is called when entering the bitwiseAndExpression production.
	EnterBitwiseAndExpression(c *BitwiseAndExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterFunctionCallExpression is called when entering the functionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterIndexExpression is called when entering the indexExpression production.
	EnterIndexExpression(c *IndexExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterLiteralExpression is called when entering the literalExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterTernaryExpression is called when entering the ternaryExpression production.
	EnterTernaryExpression(c *TernaryExpressionContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterTypeReference is called when entering the typeReference production.
	EnterTypeReference(c *TypeReferenceContext)

	// EnterTypeInstantiation is called when entering the typeInstantiation production.
	EnterTypeInstantiation(c *TypeInstantiationContext)

	// EnterBuiltinType is called when entering the builtinType production.
	EnterBuiltinType(c *BuiltinTypeContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterTypeArgument is called when entering the typeArgument production.
	EnterTypeArgument(c *TypeArgumentContext)

	// EnterDynamicLengthArgument is called when entering the dynamicLengthArgument production.
	EnterDynamicLengthArgument(c *DynamicLengthArgumentContext)

	// EnterIntType is called when entering the intType production.
	EnterIntType(c *IntTypeContext)

	// EnterVarintType is called when entering the varintType production.
	EnterVarintType(c *VarintTypeContext)

	// EnterFixedBitFieldType is called when entering the fixedBitFieldType production.
	EnterFixedBitFieldType(c *FixedBitFieldTypeContext)

	// EnterDynamicBitFieldType is called when entering the dynamicBitFieldType production.
	EnterDynamicBitFieldType(c *DynamicBitFieldTypeContext)

	// EnterBoolType is called when entering the boolType production.
	EnterBoolType(c *BoolTypeContext)

	// EnterStringType is called when entering the stringType production.
	EnterStringType(c *StringTypeContext)

	// EnterFloatType is called when entering the floatType production.
	EnterFloatType(c *FloatTypeContext)

	// EnterExternType is called when entering the externType production.
	EnterExternType(c *ExternTypeContext)

	// ExitPackageDeclaration is called when exiting the packageDeclaration production.
	ExitPackageDeclaration(c *PackageDeclarationContext)

	// ExitPackageNameDefinition is called when exiting the packageNameDefinition production.
	ExitPackageNameDefinition(c *PackageNameDefinitionContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitLanguageDirective is called when exiting the languageDirective production.
	ExitLanguageDirective(c *LanguageDirectiveContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitSymbolDefinition is called when exiting the symbolDefinition production.
	ExitSymbolDefinition(c *SymbolDefinitionContext)

	// ExitConstDefinition is called when exiting the constDefinition production.
	ExitConstDefinition(c *ConstDefinitionContext)

	// ExitRuleGroupDefinition is called when exiting the ruleGroupDefinition production.
	ExitRuleGroupDefinition(c *RuleGroupDefinitionContext)

	// ExitRuleDefinition is called when exiting the ruleDefinition production.
	ExitRuleDefinition(c *RuleDefinitionContext)

	// ExitSubtypeDeclaration is called when exiting the subtypeDeclaration production.
	ExitSubtypeDeclaration(c *SubtypeDeclarationContext)

	// ExitStructureDeclaration is called when exiting the structureDeclaration production.
	ExitStructureDeclaration(c *StructureDeclarationContext)

	// ExitStructureFieldDefinition is called when exiting the structureFieldDefinition production.
	ExitStructureFieldDefinition(c *StructureFieldDefinitionContext)

	// ExitFieldAlignment is called when exiting the fieldAlignment production.
	ExitFieldAlignment(c *FieldAlignmentContext)

	// ExitFieldOffset is called when exiting the fieldOffset production.
	ExitFieldOffset(c *FieldOffsetContext)

	// ExitFieldTypeId is called when exiting the fieldTypeId production.
	ExitFieldTypeId(c *FieldTypeIdContext)

	// ExitFieldArrayRange is called when exiting the fieldArrayRange production.
	ExitFieldArrayRange(c *FieldArrayRangeContext)

	// ExitFieldInitializer is called when exiting the fieldInitializer production.
	ExitFieldInitializer(c *FieldInitializerContext)

	// ExitFieldOptionalClause is called when exiting the fieldOptionalClause production.
	ExitFieldOptionalClause(c *FieldOptionalClauseContext)

	// ExitFieldConstraint is called when exiting the fieldConstraint production.
	ExitFieldConstraint(c *FieldConstraintContext)

	// ExitChoiceDeclaration is called when exiting the choiceDeclaration production.
	ExitChoiceDeclaration(c *ChoiceDeclarationContext)

	// ExitChoiceCases is called when exiting the choiceCases production.
	ExitChoiceCases(c *ChoiceCasesContext)

	// ExitChoiceCase is called when exiting the choiceCase production.
	ExitChoiceCase(c *ChoiceCaseContext)

	// ExitChoiceDefault is called when exiting the choiceDefault production.
	ExitChoiceDefault(c *ChoiceDefaultContext)

	// ExitChoiceFieldDefinition is called when exiting the choiceFieldDefinition production.
	ExitChoiceFieldDefinition(c *ChoiceFieldDefinitionContext)

	// ExitUnionDeclaration is called when exiting the unionDeclaration production.
	ExitUnionDeclaration(c *UnionDeclarationContext)

	// ExitUnionFieldDefinition is called when exiting the unionFieldDefinition production.
	ExitUnionFieldDefinition(c *UnionFieldDefinitionContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitEnumItem is called when exiting the enumItem production.
	ExitEnumItem(c *EnumItemContext)

	// ExitBitmaskDeclaration is called when exiting the bitmaskDeclaration production.
	ExitBitmaskDeclaration(c *BitmaskDeclarationContext)

	// ExitBitmaskValue is called when exiting the bitmaskValue production.
	ExitBitmaskValue(c *BitmaskValueContext)

	// ExitSqlTableDeclaration is called when exiting the sqlTableDeclaration production.
	ExitSqlTableDeclaration(c *SqlTableDeclarationContext)

	// ExitSqlTableFieldDefinition is called when exiting the sqlTableFieldDefinition production.
	ExitSqlTableFieldDefinition(c *SqlTableFieldDefinitionContext)

	// ExitSqlConstraintDefinition is called when exiting the sqlConstraintDefinition production.
	ExitSqlConstraintDefinition(c *SqlConstraintDefinitionContext)

	// ExitSqlConstraint is called when exiting the sqlConstraint production.
	ExitSqlConstraint(c *SqlConstraintContext)

	// ExitSqlWithoutRowId is called when exiting the sqlWithoutRowId production.
	ExitSqlWithoutRowId(c *SqlWithoutRowIdContext)

	// ExitSqlDatabaseDefinition is called when exiting the sqlDatabaseDefinition production.
	ExitSqlDatabaseDefinition(c *SqlDatabaseDefinitionContext)

	// ExitSqlDatabaseFieldDefinition is called when exiting the sqlDatabaseFieldDefinition production.
	ExitSqlDatabaseFieldDefinition(c *SqlDatabaseFieldDefinitionContext)

	// ExitServiceDefinition is called when exiting the serviceDefinition production.
	ExitServiceDefinition(c *ServiceDefinitionContext)

	// ExitServiceMethodDefinition is called when exiting the serviceMethodDefinition production.
	ExitServiceMethodDefinition(c *ServiceMethodDefinitionContext)

	// ExitPubsubDefinition is called when exiting the pubsubDefinition production.
	ExitPubsubDefinition(c *PubsubDefinitionContext)

	// ExitPubsubMessageDefinition is called when exiting the pubsubMessageDefinition production.
	ExitPubsubMessageDefinition(c *PubsubMessageDefinitionContext)

	// ExitTopicDefinition is called when exiting the topicDefinition production.
	ExitTopicDefinition(c *TopicDefinitionContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitParameterDefinition is called when exiting the parameterDefinition production.
	ExitParameterDefinition(c *ParameterDefinitionContext)

	// ExitTemplateParameters is called when exiting the templateParameters production.
	ExitTemplateParameters(c *TemplateParametersContext)

	// ExitTemplateArguments is called when exiting the templateArguments production.
	ExitTemplateArguments(c *TemplateArgumentsContext)

	// ExitTemplateArgument is called when exiting the templateArgument production.
	ExitTemplateArgument(c *TemplateArgumentContext)

	// ExitInstantiateDeclaration is called when exiting the instantiateDeclaration production.
	ExitInstantiateDeclaration(c *InstantiateDeclarationContext)

	// ExitBitwiseXorExpression is called when exiting the bitwiseXorExpression production.
	ExitBitwiseXorExpression(c *BitwiseXorExpressionContext)

	// ExitDotExpression is called when exiting the dotExpression production.
	ExitDotExpression(c *DotExpressionContext)

	// ExitValueofExpression is called when exiting the valueofExpression production.
	ExitValueofExpression(c *ValueofExpressionContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitArrayExpression is called when exiting the arrayExpression production.
	ExitArrayExpression(c *ArrayExpressionContext)

	// ExitNumbitsExpression is called when exiting the numbitsExpression production.
	ExitNumbitsExpression(c *NumbitsExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitLengthofExpression is called when exiting the lengthofExpression production.
	ExitLengthofExpression(c *LengthofExpressionContext)

	// ExitIdentifierExpression is called when exiting the identifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitBitwiseOrExpression is called when exiting the bitwiseOrExpression production.
	ExitBitwiseOrExpression(c *BitwiseOrExpressionContext)

	// ExitParenthesizedExpression is called when exiting the parenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitBitwiseAndExpression is called when exiting the bitwiseAndExpression production.
	ExitBitwiseAndExpression(c *BitwiseAndExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitFunctionCallExpression is called when exiting the functionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitIndexExpression is called when exiting the indexExpression production.
	ExitIndexExpression(c *IndexExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitLiteralExpression is called when exiting the literalExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitTernaryExpression is called when exiting the ternaryExpression production.
	ExitTernaryExpression(c *TernaryExpressionContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitTypeReference is called when exiting the typeReference production.
	ExitTypeReference(c *TypeReferenceContext)

	// ExitTypeInstantiation is called when exiting the typeInstantiation production.
	ExitTypeInstantiation(c *TypeInstantiationContext)

	// ExitBuiltinType is called when exiting the builtinType production.
	ExitBuiltinType(c *BuiltinTypeContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitTypeArgument is called when exiting the typeArgument production.
	ExitTypeArgument(c *TypeArgumentContext)

	// ExitDynamicLengthArgument is called when exiting the dynamicLengthArgument production.
	ExitDynamicLengthArgument(c *DynamicLengthArgumentContext)

	// ExitIntType is called when exiting the intType production.
	ExitIntType(c *IntTypeContext)

	// ExitVarintType is called when exiting the varintType production.
	ExitVarintType(c *VarintTypeContext)

	// ExitFixedBitFieldType is called when exiting the fixedBitFieldType production.
	ExitFixedBitFieldType(c *FixedBitFieldTypeContext)

	// ExitDynamicBitFieldType is called when exiting the dynamicBitFieldType production.
	ExitDynamicBitFieldType(c *DynamicBitFieldTypeContext)

	// ExitBoolType is called when exiting the boolType production.
	ExitBoolType(c *BoolTypeContext)

	// ExitStringType is called when exiting the stringType production.
	ExitStringType(c *StringTypeContext)

	// ExitFloatType is called when exiting the floatType production.
	ExitFloatType(c *FloatTypeContext)

	// ExitExternType is called when exiting the externType production.
	ExitExternType(c *ExternTypeContext)
}
