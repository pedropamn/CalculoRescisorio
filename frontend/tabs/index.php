<!DOCTYPE html>
<html lang="en" >
<head>
  <meta charset="UTF-8">
  <title>Featured Tabs</title>
  <link rel='stylesheet' href='../assets/css/ribbon.css'> <!-- https://codepen.io/magnusriga/pen/aKopeG -->
  <link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css'>
<link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css'>
<link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Source+Sans+Pro:400,200,300,700'><link rel="stylesheet" href="./style.css">

</head>
<body>
<!-- partial:index.partial.html -->
<div class="container"> 
<section id="fancyTabWidget" class="tabs t-tabs">
        <ul class="nav nav-tabs fancyTabs" role="tablist">
        
                    <li class="tab fancyTab active">
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>	
                        <a id="tab0" href="#tabBody0" role="tab" aria-controls="tabBody0" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-fire"></span><span class="hidden-xs">Rescisão</span></a>
                    	<div class="whiteBlock"></div>
                    </li>
                    
                    <li class="tab fancyTab">
					  <div class="ribbon blue"><span>Em breve</span></div>
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>
                        <a id="tab1" href="#tabBody1" role="tab" aria-controls="tabBody1" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-firefox"></span><span class="hidden-xs">Férias</span></a>
                        <div class="whiteBlock"></div>
                    </li>
                    
                    <li class="tab fancyTab">
					<div class="ribbon blue"><span>Em breve</span></div>
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>
                        <a id="tab2" href="#tabBody2" role="tab" aria-controls="tabBody2" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-envira"></span><span class="hidden-xs">Décimo Terceiro</span></a>
                        <div class="whiteBlock"></div>
                    </li>
                    
                    <li class="tab fancyTab">
					<div class="ribbon blue"><span>Em breve</span></div>
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>
                        <a id="tab3" href="#tabBody3" role="tab" aria-controls="tabBody3" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-mortar-board"></span><span class="hidden-xs">FGTS</span></a>
                        <div class="whiteBlock"></div>
                    </li> 
                         
                    <li class="tab fancyTab">
					<div class="ribbon blue"><span>Em breve</span></div>
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>
                        <a id="tab4" href="#tabBody4" role="tab" aria-controls="tabBody4" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-stack-overflow"></span><span class="hidden-xs">Seguro Desemprego</span></a>
                        <div class="whiteBlock"></div>
                    </li>
                    
                    <li class="tab fancyTab">
					<div class="ribbon blue"><span>Em breve</span></div>
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>
                        <a id="tab5" href="#tabBody5" role="tab" aria-controls="tabBody5" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-question-circle"></span><span class="hidden-xs">Adicional Noturno</span></a>
                        <div class="whiteBlock"></div>
                    </li>
        </ul>
        <div id="myTabContent" class="tab-content fancyTabContent" aria-live="polite">
                    <div class="tab-pane  fade active in" id="tabBody0" role="tabpanel" aria-labelledby="tab0" aria-hidden="false" tabindex="0">
                        <div>
                        	<div class="row">
                            	
                                <div class="col-lg-12">
									<form action="forms/send.php" method="post" class="php-email-form aos-init aos-animate" data-aos="fade-up" data-aos-delay="200">
										<div class="row gy-4">

											<!-- Input para Salário Mensal -->
											<div class="col-md-6">
												<label for="salarioMensal" class="form-label">Salário Mensal</label>
												<input type="number" class="form-control" id="salarioMensal" name="salarioMensal" placeholder="Ex.: 3000.00" step="0.01" required>
											</div>

											<!-- Input para Dias Trabalhados no Mês -->
											<div class="col-md-6">
												<label for="diasTrabalhados" class="form-label">Dias Trabalhados no Mês</label>
												<input type="number" class="form-control" id="diasTrabalhados" name="diasTrabalhados" placeholder="Ex.: 20" required>
											</div>

											<!-- Input para Anos de Empresa -->
											<div class="col-md-6">
												<label for="anosDeEmpresa" class="form-label">Anos de Empresa</label>
												<input type="number" class="form-control" id="anosDeEmpresa" name="anosDeEmpresa" placeholder="Ex.: 5" required>
											</div>

											<!-- Input para Meses Trabalhados Esse Ano -->
											<div class="col-md-6">
												<label for="mesesTrabalhados" class="form-label">Meses Trabalhados Esse Ano</label>
												<input type="number" class="form-control" id="mesesTrabalhados" name="mesesTrabalhados" placeholder="Ex.: 8" required>
											</div>

											<div class="col-md-12 text-center">
												<div class="loading" style="display:none;">Carregando</div>
												<div class="error-message"></div>
												<div class="sent-message" style="display:none;">Sua mensagem foi enviada. Obrigado!</div>
												
												<br>
												<button class="btn btn-success btn-lg"type="submit">Enviar</button>
											</div>

										</div>
									</form>
								</div>
                                
                            </div>
                        </div>
                    </div>
                    <div class="tab-pane  fade" id="tabBody1" role="tabpanel" aria-labelledby="tab1" aria-hidden="true" tabindex="0">
                        <div class="row">
                            	
                                <div class="col-md-12">
                        			<h2>Em breve</h2>
                                    <p>Em breve</p>
                                   
                                </div>
                            </div>
                    </div>
                    <div class="tab-pane  fade" id="tabBody2" role="tabpanel" aria-labelledby="tab2" aria-hidden="true" tabindex="0">
                        <div class="row">
							<div class="col-md-12">
								<h2>Em breve</h2>
								<p>Em breve</p>                                  
							</div>
						</div>
                    </div>
                    <div class="tab-pane  fade" id="tabBody3" role="tabpanel" aria-labelledby="tab3" aria-hidden="true" tabindex="0">
						 <div class="row">
							<div class="col-md-12">
								<h2>Em breve</h2>
								<p>Em breve</p>                                  
							</div>
						</div>
                    </div>
                    <div class="tab-pane  fade" id="tabBody4" role="tabpanel" aria-labelledby="tab4" aria-hidden="true" tabindex="0">
                     <div class="row">
							<div class="col-md-12">
								<h2>Em breve</h2>
								<p>Em breve</p>                                  
							</div>
						</div>
                    </div>
                    <div class="tab-pane  fade" id="tabBody5" role="tabpanel" aria-labelledby="tab5" aria-hidden="true" tabindex="0">
					
                     <div class="row">
							<div class="col-md-12">
								<h2>Em breve</h2>
								<p>Em breve</p>                                  
							</div>
						</div>
                    </div>
        </div>

    </section>
</div>
<!-- partial -->
  <script src='//cdnjs.cloudflare.com/ajax/libs/jquery/2.1.3/jquery.min.js'></script>
<script src='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js'></script><script  src="./script.js"></script>

</body>
</html>
