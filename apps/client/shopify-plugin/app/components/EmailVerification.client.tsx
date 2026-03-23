import {Banner, BlockStack, Card} from '@shopify/polaris';
import {useCallback, useEffect, useMemo, useState} from 'react';
import {checkEmailVerificationStatus, sendVerificationEmail} from '../service/common.service';
import type {TpShopStorageInterface} from '../interfaces/storage.interface';
import {logout} from '../utils/auth.utils';
import {useAppBridge} from '@shopify/app-bridge-react';

type ButtonState = 'resend' | 'check';

export default function EmailVerification() {
	const tpShop: TpShopStorageInterface = useMemo(
		() => JSON.parse(sessionStorage.getItem('tp_shop')!),
		[],
	);
	const [isVerified, setIsVerified] = useState<boolean | null>(null);
	const [isLoading, setIsLoading] = useState(true);
	const [isSending, setIsSending] = useState(false);
	const [buttonState, setButtonState] = useState<ButtonState>('resend');
	const appBridge = useAppBridge();

	const checkStatus = useCallback(() => {
		setIsLoading(true);
		checkEmailVerificationStatus()
			.then((response) => {
				if (response.status === 401) {
					logout(tpShop.shopUuid);
					return;
				}

				if (response.status === 200 && response.data) {
					if (response.data.isVerified) {
						// User is now verified, hide the component
						setIsVerified(true);
						appBridge.toast.show('Email verified successfully!', {
							duration: 5000,
						});
					} else {
						// Not verified yet, go back to resend state
						setIsVerified(false);
						setButtonState('resend');
						appBridge.toast.show('Email not verified yet. Please check your inbox.', {
							duration: 5000,
							isError: true,
						});
					}
				}
				setIsLoading(false);
			})
			.catch((error: Error) => {
				console.error('Error checking verification status:', error);
				setButtonState('resend');
				setIsLoading(false);
			});
	}, [tpShop, appBridge]);

	useEffect(() => {
		if (!tpShop) {
			logout(window.location.hostname);
			return;
		}

		// Initial check on mount
		setIsLoading(true);
		checkEmailVerificationStatus()
			.then((response) => {
				if (response.status === 401) {
					logout(tpShop.shopUuid);
					return;
				}

				if (response.status !== 200 || !response.data) {
					setIsLoading(false);
					return;
				}

				setIsVerified(response.data.isVerified);
				setIsLoading(false);
			})
			.catch((error: Error) => {
				console.error('Error fetching verification status:', error);
				setIsLoading(false);
			});
	}, [tpShop]);

	const handleResendEmail = useCallback(() => {
		setIsSending(true);
		sendVerificationEmail()
			.then((response) => {
				if (response.status === 401) {
					logout(tpShop.shopUuid);
					return;
				}

				if (response.status === 200) {
					appBridge.toast.show('Verification email sent successfully', {
						duration: 5000,
					});
					// Change button to "Check verification status"
					setButtonState('check');
				} else {
					appBridge.toast.show('Failed to send verification email. Please try again.', {
						duration: 5000,
						isError: true,
					});
				}
				setIsSending(false);
			})
			.catch((error: Error) => {
				console.error('Error sending verification email:', error);
				appBridge.toast.show('Failed to send verification email. Please try again.', {
					duration: 5000,
					isError: true,
				});
				setIsSending(false);
			});
	}, [tpShop, appBridge]);

	const handleButtonClick = useCallback(() => {
		if (buttonState === 'resend') {
			handleResendEmail();
		} else {
			checkStatus();
		}
	}, [buttonState, handleResendEmail, checkStatus]);

	// If loading or user is verified, don't show anything
	if (isLoading || isVerified === true) {
		return null;
	}

	const buttonText = buttonState === 'resend' ? 'Resend Email' : 'Check Verification';

	return (
		<Card>
			<BlockStack gap="400">
				<Banner
					tone="warning"
					title="Email Verification Required"
					action={{
						content: buttonText,
						onAction: handleButtonClick,
						loading: isSending || isLoading,
					}}
				>
					<p>
						Please verify your email address to ensure account security and receive important
						notifications.
					</p>
				</Banner>
			</BlockStack>
		</Card>
	);
}
