# 🤝 Contributing Guide

Thank you for your interest in contributing to the Debt Tracker project! This guide will help you get started with contributing effectively.

## 🎯 Project Vision

We're building a comprehensive personal finance management tool that combines manual debt tracking with automated bill monitoring through bank integration. Your contributions help make financial management accessible and intelligent for everyone.

## 🌟 Ways to Contribute

### 🔥 High Priority Areas
- **🏦 Bank Integration**: Kiwi API implementation and data processing
- **📱 Mobile Development**: React Native/Flutter applications
- **🧪 Testing**: Unit tests, integration tests, and end-to-end testing
- **🔒 Security**: Security audits, improvements, and best practices
- **📊 Analytics**: Advanced financial insights and visualizations

### 📈 Medium Priority Areas
- **🎨 UI/UX**: Design improvements and user experience enhancements
- **📖 Documentation**: Tutorials, guides, and API documentation
- **🇿🇦 South African Localization**: ZAR currency, local banking integration, SA-specific features
- **⚡ Performance**: Optimization and scalability improvements
- **🔗 Integrations**: South African financial service integrations

### 🌱 Good for Beginners
- **🐛 Bug Fixes**: Small bug fixes and improvements
- **📝 Documentation**: README improvements, code comments
- **🎨 Frontend**: CSS improvements and responsive design
- **🧹 Code Quality**: Refactoring and code cleanup
- **🧪 Test Writing**: Writing tests for existing features

## 📋 Prerequisites

### System Requirements
- **Ubuntu 20.04+ or WSL with Ubuntu** (required)
- **Go 1.19+** for backend development
- **Flutter SDK** for mobile development
- **Android Studio** for Android development and emulation
- **Git** for version control

### Skills Helpful for Contributing
- **Backend**: Go, REST APIs, SQL, JWT authentication
- **Frontend**: HTML/CSS/JavaScript, Alpine.js, responsive design
- **Mobile**: Flutter, Dart programming language
- **Database**: SQLite, PostgreSQL (for production features)
- **South African Context**: Local banking systems, ZAR currency, SA financial regulations

## 🚀 Getting Started

### 1. Fork and Clone

1. **Fork the repository** on GitHub
2. **Clone your fork**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/debt-tracker-app.git
   cd debt-tracker-app
   ```

3. **Add upstream remote**:
   ```bash
   git remote add upstream https://github.com/MAIN_USERNAME/debt-tracker-app.git
   ```

### 2. Set Up Development Environment

Follow the complete setup guide in [SETUP.md](./SETUP.md).

### 3. Create a Feature Branch

```bash
# Update your main branch
git checkout main
git pull upstream main

# Create a new feature branch
git checkout -b feature/your-feature-name

# Examples:
# git checkout -b feature/kiwi-api-integration
# git checkout -b fix/transaction-validation-bug
# git checkout -b docs/api-documentation-update
```

## 💻 Development Workflow

### 1. Code Standards

**Go Backend Standards:**
```go
// Use meaningful variable names
func CreateDebt(userID int, request CreateDebtRequest) (*Debt, error) {
    // Validate input
    if request.Amount <= 0 {
        return nil, errors.New("amount must be positive")
    }
    
    // Use proper error handling
    debt, err := db.CreateDebt(userID, request)
    if err != nil {
        return nil, fmt.Errorf("failed to create debt: %w", err)
    }
    
    return debt, nil
}
```

**Frontend Standards:**
```javascript
// Use descriptive function names
async function createDebtForContact(contactId, amount, direction, description) {
    try {
        const response = await apiCall('/debts', {
            method: 'POST',
            body: JSON.stringify({
                contact_id: contactId,
                amount: parseFloat(amount),
                direction,
                description
            })
        });
        return response;
    } catch (error) {
        console.error('Failed to create debt:', error);
        throw error;
    }
}
```

### 2. Coding Guidelines

**General Principles:**
- **Clean Code**: Write self-documenting code with clear variable names
- **Error Handling**: Always handle errors gracefully
- **Security First**: Validate all inputs and sanitize outputs
- **Performance**: Consider performance implications of your changes
- **Documentation**: Comment complex logic and add docstrings

**Go-Specific Guidelines:**
- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` for formatting
- Run `go vet` for static analysis
- Use meaningful package names
- Handle all errors explicitly

**Frontend Guidelines:**
- Use semantic HTML elements
- Ensure responsive design for mobile devices
- Add appropriate ARIA labels for accessibility
- Optimize images and assets
- Use progressive enhancement principles

### 3. Commit Message Format

Use the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

**Examples:**
```
feat(api): add Kiwi API integration for bank transactions

- Implement OAuth2 flow for bank authentication
- Add transaction fetching and parsing
- Create categorization service for expenses

Closes #123
```

```
fix(frontend): resolve transaction modal validation issue

The transaction amount field was not properly validating
negative numbers, causing API errors.

Fixes #456
```

```
docs(setup): update Ubuntu installation instructions

Added WSL-specific setup steps and troubleshooting
section for Windows users.
```

## 🧪 Testing Guidelines

### Running Tests

```bash
# Backend tests
cd backend
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Frontend tests (when implemented)
cd frontend
npm test
```

### Writing Tests

**Backend Test Example:**
```go
func TestCreateDebt(t *testing.T) {
    // Setup
    db := setupTestDB(t)
    defer teardownTestDB(db)
    
    userID := createTestUser(t, db)
    contactID := createTestContact(t, db, userID)
    
    // Test case
    request := CreateDebtRequest{
        ContactID:   contactID,
        Amount:      50.00,
        Direction:   "owe_to",
        Description: "Test debt",
    }
    
    debt, err := CreateDebt(userID, request)
    
    // Assertions
    assert.NoError(t, err)
    assert.Equal(t, 50.00, debt.Amount)
    assert.Equal(t, "owe_to", debt.Direction)
    assert.Equal(t, "Test debt", debt.Description)
}
```

**Frontend Test Example:**
```javascript
describe('Debt Creation', () => {
    test('should create debt with valid data', async () => {
        const mockApiCall = jest.fn().mockResolvedValue({
            id: 1,
            amount: 50.00,
            direction: 'owe_to'
        });
        
        const result = await createDebt(1, 50.00, 'owe_to', 'Test debt');
        
        expect(mockApiCall).toHaveBeenCalledWith('/debts', {
            method: 'POST',
            body: JSON.stringify({
                contact_id: 1,
                amount: 50.00,
                direction: 'owe_to',
                description: 'Test debt'
            })
        });
        
        expect(result.amount).toBe(50.00);
    });
});
```

### Test Coverage Requirements

- **New Features**: Must include comprehensive tests
- **Bug Fixes**: Must include regression tests
- **Critical Paths**: Authentication, data validation, financial calculations
- **Target Coverage**: Aim for 80%+ code coverage

## 📖 Documentation Standards

### Code Documentation

**Go Documentation:**
```go
// CreateDebt creates a new debt record for the specified user and contact.
// It validates the input data and returns the created debt or an error.
//
// Parameters:
//   - userID: The ID of the user creating the debt
//   - request: The debt creation request containing contact_id, amount, direction, and description
//
// Returns:
//   - *Debt: The created debt record
//   - error: Any error that occurred during creation
func CreateDebt(userID int, request CreateDebtRequest) (*Debt, error) {
    // Implementation...
}
```

**API Documentation:**
```go
// @Summary Create a new debt
// @Description Create a new debt record for a contact
// @Tags debts
// @Accept json
// @Produce json
// @Param debt body CreateDebtRequest true "Debt data"
// @Success 201 {object} Debt
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /debts [post]
func (h *DebtHandler) CreateDebt(c *gin.Context) {
    // Implementation...
}
```

### README and Guide Updates

When adding new features:
1. Update relevant README files
2. Add usage examples
3. Update API documentation
4. Include troubleshooting tips
5. Add configuration instructions

## 🔒 Security Guidelines

### Security Checklist

**Before Submitting:**
- [ ] All user inputs are validated and sanitized
- [ ] SQL queries use parameterized statements
- [ ] Sensitive data is not logged
- [ ] Authentication is properly implemented
- [ ] Authorization checks are in place
- [ ] Error messages don't leak sensitive information

**Security Best Practices:**
- Never commit secrets or API keys
- Use environment variables for configuration
- Implement rate limiting for API endpoints
- Validate file uploads and limit file sizes
- Use HTTPS in production
- Implement proper session management

### Reporting Security Issues

If you discover a security vulnerability:
1. **DO NOT** create a public GitHub issue
2. Email the maintainers directly: security@debt-tracker.app
3. Include detailed steps to reproduce
4. Allow reasonable time for response and fixes

## 🐛 Bug Reports

### Before Reporting a Bug

1. **Search existing issues** to avoid duplicates
2. **Update to the latest version** and test again
3. **Check the documentation** for known limitations
4. **Test in a clean environment** to rule out local issues

### Bug Report Template

```markdown
**Bug Description**
A clear description of what the bug is.

**Steps to Reproduce**
1. Go to '...'
2. Click on '...'
3. Scroll down to '...'
4. See error

**Expected Behavior**
What you expected to happen.

**Actual Behavior**
What actually happened.

**Environment**
- OS: Ubuntu 22.04
- Go Version: 1.21.0
- Browser: Chrome 118.0
- API Version: v1.0.0

**Screenshots**
If applicable, add screenshots.

**Additional Context**
Any other context about the problem.
```

## ✨ Feature Requests

### Feature Request Template

```markdown
**Feature Description**
A clear description of the feature you'd like to see.

**Problem Statement**
What problem does this feature solve? Who would benefit?

**Proposed Solution**
How would you like this feature to work?

**Alternative Solutions**
Any alternative approaches you've considered.

**Additional Context**
Mockups, examples, or references to similar features.

**Priority Level**
- [ ] Low (nice to have)
- [ ] Medium (would improve workflow)
- [ ] High (blocking current workflow)
- [ ] Critical (security or data integrity issue)
```

## 📝 Pull Request Process

### Before Creating a Pull Request

1. **Ensure your branch is up to date:**
   ```bash
   git checkout main
   git pull upstream main
   git checkout your-feature-branch
   git rebase main
   ```

2. **Run the full test suite:**
   ```bash
   # Backend tests
   cd backend && go test ./...
   
   # Frontend tests (when available)
   cd frontend && npm test
   
   # Integration tests
   ./scripts/test-integration.sh
   ```

3. **Format your code:**
   ```bash
   # Go formatting
   cd backend && go fmt ./...
   
   # Frontend formatting
   cd frontend && prettier --write .
   ```

### Pull Request Template

```markdown
## Description
Brief description of the changes in this PR.

## Type of Change
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Performance improvement
- [ ] Code refactoring

## Related Issues
Closes #123
Fixes #456
Related to #789

## Testing
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Manual testing completed
- [ ] New tests added (if applicable)

## Screenshots
(If applicable, especially for UI changes)

## Checklist
- [ ] My code follows the project's coding standards
- [ ] I have performed a self-review of my code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
```

### Review Process

1. **Automated Checks**: CI/CD pipeline runs tests and checks
2. **Code Review**: At least one maintainer reviews the code
3. **Testing**: Feature is tested in a staging environment
4. **Documentation**: Ensure documentation is updated
5. **Merge**: Approved PRs are merged to main branch

### Review Criteria

**Code Quality:**
- Follows coding standards and conventions
- Includes appropriate error handling
- Has adequate test coverage
- Is well-documented

**Functionality:**
- Solves the intended problem
- Doesn't break existing functionality
- Performs well under expected load
- Follows security best practices

**User Experience:**
- Intuitive and user-friendly
- Accessible to users with disabilities
- Works across different devices and browsers
- Provides clear feedback to users

## 🏗️ Development Areas

### Backend Development (Go)

**Current Tech Stack:**
- Go 1.21+ with Gin framework
- SQLite database (PostgreSQL for production)
- JWT authentication
- RESTful API design

**Key Areas for Contribution:**
- Kiwi API integration for bank data
- Advanced analytics and reporting
- Performance optimization
- Security enhancements
- Database migrations and optimization

**Getting Started:**
1. Read the existing API code in `internal/handlers/`
2. Understand the database schema in `internal/database/`
3. Study the authentication middleware
4. Check open issues tagged with `backend`

### Frontend Development

**Current Tech Stack:**
- Vanilla HTML/CSS/JavaScript
- Alpine.js for reactivity
- Chart.js for visualizations
- Responsive CSS design
- Bootstrap-inspired components

**Key Areas for Contribution:**
- Mobile-responsive improvements for South African users
- ZAR currency display components
- South African banking UI patterns
- Data visualization enhancements
- Progressive Web App features

**South African UI Considerations:**
- Mobile-first design (high mobile usage in SA)
- Low-bandwidth optimization for data-conscious users
- Offline capability for areas with poor connectivity
- Touch-friendly interfaces for mobile devices
- Local color schemes and cultural preferences

**Getting Started:**
1. Review the main dashboard in `frontend/index.html`
2. Understand the Alpine.js reactive patterns
3. Test the interface on various mobile devices
4. Research South African mobile usage patterns
5. Check open issues tagged with `frontend` or `ui-sa`

### Mobile Development

**Target Platform: Flutter Mobile App 📱**
- Flutter SDK for cross-platform development
- Native Android and iOS applications
- Dart programming language
- South African market focus
- Integration with local payment systems

**Key Areas for Contribution:**
- Flutter app development from scratch
- Native device feature integration (camera, notifications)
- Offline-first architecture for reliable operation
- South African banking integration
- ZAR currency handling and formatting

**Getting Started:**
1. Install Flutter SDK and Android Studio
2. Set up Flutter development environment on Ubuntu
3. Study the existing API integration patterns
4. Design mobile-first user flows for South African users
5. Check open issues tagged with `flutter` or `mobile`

**Flutter Setup Commands:**
```bash
# Install Flutter
git clone https://github.com/flutter/flutter.git -b stable
export PATH="$PATH:`pwd`/flutter/bin"

# Verify installation
flutter doctor

# Create Flutter project
flutter create debt_tracker_mobile
cd debt_tracker_mobile

# Run on Android emulator
flutter run
```

### Bank Integration (High Priority - South Africa Focus)

**Kiwi API Integration for South African Banks:**
- OAuth2 authentication flow with SA banks
- Transaction data fetching from major SA banks (FNB, Standard Bank, ABSA, Nedbank)
- ZAR currency handling and formatting
- South African transaction categorization
- Compliance with SA financial regulations

**Key Areas for Contribution:**
- API client implementation for South African banking
- ZAR transaction processing pipelines
- SA-specific transaction categorization (municipal bills, school fees, etc.)
- User consent flows compliant with SA privacy laws
- Integration testing with SA banking systems

**South African Banking Context:**
- Major banks: First National Bank, Standard Bank, ABSA, Nedbank, Capitec
- Common transaction types: EFT, debit orders, municipal payments
- Local payment methods: SnapScan, Zapper, Instant EFT
- Regulatory compliance: POPI Act, SARB regulations

**Getting Started:**
1. Study Kiwi API documentation for South African banks
2. Understand South African banking transaction formats
3. Research POPI Act compliance requirements
4. Design secure data handling for ZAR transactions
5. Check open issues tagged with `bank-integration` or `south-africa`

### Testing and Quality Assurance

**Testing Priorities:**
- Unit tests for backend services
- Integration tests for API endpoints
- Frontend component testing
- End-to-end user workflow testing
- Performance and load testing

**Key Areas for Contribution:**
- Test framework setup
- Test case development
- Automated testing pipelines
- Performance benchmarking
- Security testing

**Getting Started:**
1. Study existing test patterns
2. Set up testing environments
3. Write tests for untested code
4. Check open issues tagged with `testing`

## 🌍 South African Localization

### Currency and Financial Context

**ZAR Currency Support:**
- South African Rand (ZAR) as primary currency
- Proper currency formatting (R 1,234.56)
- Cents handling and rounding rules
- Historical exchange rate tracking (for future features)

**South African Financial Landscape:**
- Major banks and their transaction formats
- Common bill types (municipal, DSTV, Eskom, etc.)
- Local payment methods and terminology
- School fees, levies, and subscription services
- South African financial calendar (tax year, payment cycles)

**Localization Areas:**
- Currency display and calculations
- Date formatting (DD/MM/YYYY South African standard)
- Banking terminology and transaction types
- Local contact number formats (+27...)
- South African English language nuances

**Getting Started:**
1. Study South African banking transaction patterns
2. Research common SA bill categories and merchants
3. Understand local payment preferences and terminology
4. Check open issues tagged with `localization` or `south-africa`

### South African Market Focus

**Target Users:**
- Young professionals managing personal finances
- Students tracking expenses and debts
- Small business owners monitoring cash flow
- Families managing household budgets

**Local Use Cases:**
- University fees and accommodation costs
- Stokvels and group savings management
- Municipal bill tracking and payments
- Retail account and credit management
- Domestic worker and service provider payments

## 🚀 Deployment and DevOps

**Key Areas:**
- Docker containerization
- CI/CD pipeline setup
- Production deployment automation
- Monitoring and logging
- Backup and disaster recovery

**Getting Started:**
1. Create Dockerfile for the application
2. Set up GitHub Actions workflows
3. Design production architecture
4. Check open issues tagged with `devops`

## 📞 Getting Help

### Communication Channels

- **GitHub Issues**: For bug reports and feature requests
- **GitHub Discussions**: For questions and general discussion
- **Discord/Slack**: For real-time collaboration (links coming soon)
- **Email**: For security issues and private matters

### Mentorship

New contributors can get help from experienced maintainers:
- Code review and feedback
- Architecture guidance
- Best practices advice
- Career development support

### Learning Resources

**Go Development:**
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

**Flutter Development:**
- [Flutter Documentation](https://flutter.dev/docs)
- [Dart Language Tour](https://dart.dev/guides/language/language-tour)
- [Flutter Cookbook](https://flutter.dev/docs/cookbook)
- [Material Design for Flutter](https://flutter.dev/docs/development/ui/material)

**South African Context:**
- [South African Reserve Bank Guidelines](https://www.resbank.co.za/)
- [POPI Act Compliance](https://popia.co.za/)
- [South African Banking Association](https://www.banking.org.za/)
- [Local Payment Methods Guide](https://www.payfast.co.za/)

## 🎉 Recognition

### Contributor Recognition

We value all contributions and recognize contributors in several ways:

- **Contributors Page**: Featured on project website
- **Release Notes**: Mentioned in release announcements
- **Social Media**: Highlighted on project social media
- **Swag**: Project stickers and merchandise for significant contributions
- **References**: LinkedIn recommendations for substantial contributors

### Contribution Levels

- **🌱 First-Time Contributor**: Welcome package and mentorship
- **🔧 Regular Contributor**: Recognition in release notes
- **🌟 Core Contributor**: Voting rights on project decisions
- **👑 Maintainer**: Full repository access and responsibility

## 📄 License

By contributing to this project, you agree that your contributions will be licensed under the project's MIT License.

---

**Thank you for contributing to Debt Tracker! Your efforts help create better financial tools for everyone. 🚀**